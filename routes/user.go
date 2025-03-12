package routes

import (
	controllers "restpad/restpad-service/controllers/internals"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://127.0.0.1:5500"}
	// config.AllowOrigins = []string{"http://google.com", "http://facebook.com"}
	// config.AllowAllOrigins = true

	router.Use(cors.New(config))

	user := router.Group("/:id")
	{
		user.GET("/data/", controllers.GetRequestHandler())
		user.GET("/data/:cond_id/", controllers.GetRequestHandler())
		user.POST("/data/", controllers.PostRequestHandler())
		user.PUT("/data/:cond_id/", controllers.PutRequestHandler())
		user.PATCH("/data/:cond_id/", controllers.PatchRequestHandler())
		user.DELETE("/data/:cond_id/", controllers.DeleteRequestHandler())
	}
}

// package routes

// import (
// 	"fmt"
// 	"net/http"
// 	"restpad/restpad-service/controllers"

// 	"github.com/gin-gonic/gin"
// )

// func UserRouter(router *gin.Engine) {
// 	router.Use(checkAllowedOriginMiddleware())

// 	user := router.Group("/:id")
// 	{
// 		user.GET("/data/", controllers.GetRequestHandler())
// 		user.POST("/data/", controllers.PostRequestHandler())
// 		user.PATCH("/data/:cond_id/", controllers.PatchRequestHandler())
// 	}
// }

// func checkAllowedOriginMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// Get the origin from the request headers
// 		origin := c.Request.Header.Get("Origin")

// 		// Check if the origin is allowed
// 		if isValidOrigin(origin) {
// 			c.Next()
// 		} else {
// 			// If the origin is not allowed, return an error response
// 			c.JSON(http.StatusForbidden, gin.H{
// 				"error": "Invalid origin",
// 			})
// 			c.Abort()
// 		}
// 	}
// }

// func isValidOrigin(requestedOrigin string) bool {
// 	// Replace this logic with your actual check
// 	fmt.Println(requestedOrigin)
// 	return requestedOrigin == "http://google.com" || requestedOrigin == "http://localhost:8080/"
// }
