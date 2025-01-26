package server

import (
	middleware "task-tracker/api/middlewares"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Server struct {
	router *gin.Engine
	db     *pgxpool.Pool
}

func NewServer(db *pgxpool.Pool) *Server {
	// Create new gin engine with our logger
	router := gin.Default()
	return &Server{
		router: router,
		db:     db,
	}

}

// Run starts the server
func (s *Server) Run() {

	// Use recovery and your custom logger middleware
	s.router.Use(gin.Recovery())

	// CORS configuration allowing all origins with flexible customization
	s.router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, // Allow sending credentials
		MaxAge:           12 * time.Hour,

		// Custom function that can be updated to allow specific origins in the future
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	}))

	//Attaching some global middlewares
	s.router.Use(middleware.ErrorHandler())
	s.router.Use(middleware.LoggingMiddleware())

	//Map all routes to server
	s.MapRoutes()

	// Start the server at port 3000
	s.router.Run(":3000")
}
