package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	_ "github.com/knands42/lorecrafter/cmd/api/docs" // Import the docs package
	middleware2 "github.com/knands42/lorecrafter/cmd/api/middleware"
	"github.com/knands42/lorecrafter/cmd/api/routes"
	"github.com/knands42/lorecrafter/internal/config"
	"github.com/knands42/lorecrafter/internal/usecases"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/swaggo/swag"
)

// Server represents the HTTP server
type Server struct {
	router     chi.Router
	httpServer *http.Server
}

// NewServer creates a new HTTP server
func NewServer(cfg config.Config) *Server {
	router := chi.NewRouter()

	// Set up CORS
	var allowedOrigins []string
	if cfg.Profile == "dev" {
		allowedOrigins = []string{"https://lorecrafter.fly.dev", "http://localhost:8000"}
	} else {
		allowedOrigins = []string{"https://lorecrafter.fly.dev"}
	}

	corsMiddleware := cors.Handler(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	// Set up middleware
	router.Use(corsMiddleware)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(10 * time.Second))

	// Create the HTTP server
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.ServerPort),
		Handler: router,
	}

	return &Server{
		router:     router,
		httpServer: httpServer,
	}
}

// Router returns the router
func (s *Server) Router() chi.Router {
	return s.router
}

// Start starts the HTTP server
func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

// Shutdown gracefully shuts down the HTTP server
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

// SetupRoutes sets up the routes for the server
func SetupRoutes(server *Server, authHandler *routes.AuthHandler, userHandler *routes.UserHandler, authUseCase *usecases.AuthUseCase) {
	// Swagger UI
	server.router.Get("/swagger/*", httpSwagger.Handler(
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
	server.router.Route("/api", func(r chi.Router) {
		// Auth routes
		r.Route("/auth", func(r chi.Router) {
			authHandler.RegisterRoutes(r)
		})

		// Protected routes (require authentication)
		r.Group(func(r chi.Router) {
			r.Use(middleware2.AuthMiddleware(authUseCase))
			r.Get("/me", middleware2.ErrorHandlerMiddleware(userHandler.Me))
		})
	})
}
