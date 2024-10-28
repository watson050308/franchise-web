package routes

import (
	"franchise-web/app/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "franchise-web/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	routes := gin.Default()

	// CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}
	config.AllowCredentials = true

	routes.Use(cors.New(config))

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	routes.GET("/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler,
			url,
		))
	api := routes.Group("/api/v1")
	{
		api.GET("/ping", controllers.TestController)
		api.POST("/signup", controllers.CreateUser)
	}

	return routes
}
