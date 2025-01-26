package server

import (
	"task-tracker/api/repository"
	"task-tracker/api/routes"
	"task-tracker/api/usecases"
)

// function to initialise routes
func (s *Server) MapRoutes() {

	// Initialize repositories
	authRepo := repository.NewAuthRepository(s.db)
	taskRepo := repository.NewTaskRepository(s.db)

	// Initialize use cases
	authUseCase := usecases.NewAuthUseCase(authRepo)
	taskUseCase := usecases.NewTaskUseCase(taskRepo)

	// Initialize the Gin router
	routes.RegisterSwaggerRoutes(s.router)
	routes.RegisterAuthRoutes(s.router, *authUseCase)
	routes.RegisterTaxkRoutes(s.router, taskUseCase)

}
