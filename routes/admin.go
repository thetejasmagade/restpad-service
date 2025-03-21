package routes

import (
	controllers "restpad/restpad-service/controllers/app"
	"restpad/restpad-service/middlewares" // Import middleware

	"github.com/gin-gonic/gin"
)

func AdminRouter(router *gin.Engine) {
	// CORS Config
	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://localhost:3000"}

	// adminCors := cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*"}, // Only allow admin dashboard
	// 	AllowMethods:     []string{"GET", "POST", "DELETE", "OPTIONS"},
	// 	AllowHeaders:     []string{"Authorization", "Content-Type"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	// })

	// Protect routes with JWT verification
	admin := router.Group("app/api/v1")
	// admin.Use(cors.New(config))
	// admin.Use(adminCors)
	admin.Use(middlewares.VerifySupabaseJWT()) // Apply JWT middleware
	{
		admin.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{"foo": "bar"})
		})

		admin.POST("/post", controllers.PostRequestHandler())
	}
}
