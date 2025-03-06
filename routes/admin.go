package routes

import (
	controllers "restpad/restpad-service/controllers/app"

	"github.com/gin-gonic/gin"
)

func AdminRouter(router *gin.Engine) {
	admin := router.Group("/admin")
	{
		admin.POST("/post", controllers.PostRequestHandler())
	}
}
