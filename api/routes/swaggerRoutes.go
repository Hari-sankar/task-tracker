package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// RegisterTestRoutes registers the test routes
func RegisterSwaggerRoutes(router *gin.Engine) {

	swaggerRouter := router.Group("/swagger")
	{
		// Serve the Swagger JSON file
		swaggerRouter.GET("/json", func(c *gin.Context) {
			c.File("./docs/swagger.json")
		})

		// Serve the Swagger UI
		swaggerURL := ginSwagger.URL("/swagger/json")
		swaggerRouter.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerURL))

	}
}
