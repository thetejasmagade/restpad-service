package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleDBError(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"status":  http.StatusInternalServerError,
		"message": message,
	})
}
