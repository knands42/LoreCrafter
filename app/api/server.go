package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/knands42/lorecrafter/app/api/docs" // Import the docs package
	middleware2 "github.com/knands42/lorecrafter/app/api/middleware"
	"github.com/knands42/lorecrafter/app/api/routes"
	"github.com/knands42/lorecrafter/internal/config"
	"github.com/knands42/lorecrafter/internal/usecases"
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/swaggo/swag"
)

// Server represents the HTTP server
type Server struct {
	Router     chi.Router
	httpServer *http.Server

	cfg config.Config

	authUseCase          *usecases.AuthUseCase
	emailUseCase         *usecases.EmailUseCase
	aiCampaignUseCase    *usecases.AICampaignUseCase
	passwordResetUseCase *usecases.PasswordResetUseCase
	campaignUseCase      *usecases.CampaignUseCase

	authHandler     *routes.AuthHandler
	userHandler     *routes.UserHandler
	campaignHandler *routes.CampaignHandler
	repo            sqlc.Querier
}

// NewServer creates a new HTTP server
func NewServer(
	cfg config.Config,
	repo sqlc.Querier,

	authUseCase *usecases.AuthUseCase,
	campaignUseCase *usecases.CampaignUseCase,
	passwordResetUseCase *usecases.PasswordResetUseCase,
) *Server {
	router := chi.NewRouter()

	// Set up middleware
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(10 * time.Second))

	// Set up CORS middleware
	allowedOrigins := []string{"https://lorecrafter-client.vercel.app", "https://lorecrafter.fly.dev"}
	if cfg.Profile == "dev" {
		allowedOrigins = append(allowedOrigins, "http://localhost:3000", "http://localhost:8080", "http://localhost:8000")
	}

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           86400, // 24 hours
		Debug:            false,
	})

	router.Use(corsMiddleware.Handler)

	// Create the HTTP server
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.ServerPort),
		Handler: router,
	}
	server := &Server{
		Router:     router,
		httpServer: httpServer,

		cfg:  cfg,
		repo: repo,

		authUseCase:          authUseCase,
		campaignUseCase:      campaignUseCase,
		passwordResetUseCase: passwordResetUseCase,

		authHandler:     routes.NewAuthHandler(authUseCase, passwordResetUseCase),
		userHandler:     routes.NewUserHandler(),
		campaignHandler: routes.NewCampaignHandler(campaignUseCase),
	}

	server.setupRoutes()

	return server
}

// Start starts the HTTP server with optional TLS
func (s *Server) Start() {
	go func() {
		log.Printf("Starting server on port %s", s.cfg.ServerPort)
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	s.gracefulShutdown()
}

// Shutdown gracefully shuts down the HTTP server
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

// SetupRoutes sets up the routes for the server
func (s *Server) setupRoutes() {
	// Swagger UI
	s.Router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // The URL pointing to API definition
		httpSwagger.DocExpansion("none"),
		httpSwagger.UIConfig(map[string]string{
			"defaultModelsExpandDepth": "-1",
		}),
		httpSwagger.BeforeScript(`
			window.onload = function() {
				const ui = window.ui;
				const token = localStorage.getItem('jwt_token');
				if (token) {
					ui.authActions.preauthorizeApiKey('BearerAuth', token);
				}

				// Listen for login success event from the login response
				window.addEventListener('message', function(event) {
					if (event.data.type === 'login-success' && event.data.token) {
						const token = event.data.token;
						localStorage.setItem('jwt_token', token);
						ui.authActions.preauthorizeApiKey('BearerAuth', token);
					}
				});
			};
		`),
	))

	// API routes
	s.Router.Route("/api", func(r chi.Router) {
		// Auth routes
		r.Route("/auth", func(r chi.Router) {
			s.authHandler.RegisterRoutes(r)
		})

		// Protected routes (require authentication)
		r.Group(func(r chi.Router) {
			r.Use(middleware2.AuthMiddleware(s.authUseCase))
			r.Get("/me", middleware2.ErrorHandlerMiddleware(s.userHandler.Me))

			// Campaign routes
			s.campaignHandler.RegisterRoutes(r)
		})
	})
}

func (s *Server) gracefulShutdown() {
	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited properly")
}
