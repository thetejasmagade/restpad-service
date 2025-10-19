package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Supabase JWT verification middleware
func VerifySupabaseJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		allowedHosts := []string{"localhost:8080", "example.com", "api.example.com"}
		clientURL := c.Request.Host

		isValid := false
		for _, host := range allowedHosts {
			if host == clientURL {
				isValid = true
				break // Exit loop early
			}
		}

		if !isValid {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Access denied: Unauthorized client URL",
			})
			c.Abort()
			return
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		// Extract the Bearer token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		// Get Supabase JWT Secret (Find this in Supabase Dashboard → API → JWT Secret)
		supabaseSecret := os.Getenv("SUPABASE_JWT_SECRET")
		if supabaseSecret == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Supabase secret key missing"})
			c.Abort()
			return
		}

		// Parse and verify JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(supabaseSecret), nil
		})

		// Check for errors
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token | Unauthorized"})
			c.Abort()
			return
		}

		// Token is valid, proceed to the next handler
		c.Next()
	}
}
