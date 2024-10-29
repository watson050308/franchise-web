package routes

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	admin "franchise-web/app/controllers/admin"
	mid "franchise-web/app/middlewares"
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
		api.POST("/login", admin.Login)
		api.POST("/register", admin.Register)

		auth := api.Group("/")
		auth.Use(
			mid.AuthMiddleware(),
		)
		{
			auth.GET("/test", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"message": "pong",
				})
			})
		}
	}

	return routes
}
