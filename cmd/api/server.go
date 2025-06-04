package api

import (
	"context"
	"fmt"
	middleware2 "github.com/knands42/lorecrafter/cmd/api/middleware"
	"github.com/knands42/lorecrafter/cmd/api/routes"
	_ "github.com/knands42/lorecrafter/cmd/api/docs" // Import the docs package
	"github.com/knands42/lorecrafter/internal/usecases"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/swaggo/swag"
)

// Server represents the HTTP server
type Server struct {
	router     chi.Router
	httpServer *http.Server
}

// NewServer creates a new HTTP server
func NewServer(port string) *Server {
	router := chi.NewRouter()

	// Set up middleware
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(10 * time.Second))

	// Create the HTTP server
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
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
func SetupRoutes(server *Server, authHandler *routes.AuthHandler, authUseCase *usecases.AuthUseCase) {
	// Swagger UI
	server.router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // The URL pointing to API definition
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

			r.Get("/me", func(w http.ResponseWriter, r *http.Request) {
				userID := r.Context().Value(middleware2.UserIDContextKey)
				w.Write([]byte(fmt.Sprintf("Authenticated! User ID: %v", userID)))
			})
		})
	})
}
