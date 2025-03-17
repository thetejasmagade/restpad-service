package routes

import (
	controllers "restpad/restpad-service/controllers/app"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func AdminRouter(router *gin.Engine) {
	// CORS Config
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}

	router.Use(cors.New(config))

	admin := router.Group("app/api/v1")
	{
		admin.POST("/post", controllers.PostRequestHandler())
	}
}
