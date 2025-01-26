package routes

import (
	"task-tracker/api/controllers"
	middleware "task-tracker/api/middlewares"
	"task-tracker/api/usecases"

	"github.com/gin-gonic/gin"
)

// RegisterTaxkRoutes registers the task routes
func RegisterTaxkRoutes(router *gin.Engine, taskUsecase usecases.TaskUseCase) {
	taskController := controllers.NewTaskController(taskUsecase)

	taskRoutes := router.Group("/task", middleware.AuthenticateRequest)
	{
		taskRoutes.GET("/", taskController.GetAllTasks)
		taskRoutes.GET("/:taskID", taskController.GetTaskByID)
		taskRoutes.POST("/", taskController.CreateTask)
		taskRoutes.PATCH("/", taskController.UpdateTask)
		taskRoutes.DELETE("/:taskID", taskController.DeleteTask)

	}
}
