package routes

import (
	controllers "restpad/restpad-service/controllers/app"
	"restpad/restpad-service/middlewares" // Import middleware

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func AdminRouter(router *gin.Engine) {
	// CORS Config
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}

	// Protect routes with JWT verification
	admin := router.Group("app/api/v1")
	admin.Use(cors.New(config))
	admin.Use(middlewares.VerifySupabaseJWT()) // Apply JWT middleware
	{
		admin.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{"foo": "bar"})
		})

		admin.POST("/post", controllers.PostRequestHandler())
	}
}
