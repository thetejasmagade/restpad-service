package main

import (
	"log"
	"restpad/restpad-service/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {

	// Initialize Router
	router := gin.Default()
	router.RedirectTrailingSlash = false
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Your frontend origin
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Registering routes for users
	routes.UserRouter(router)

	// Registering routes for admin
	routes.AdminRouter(router)

	// err := configs.OpenDatabase()
	// if err != nil {
	// 	log.Println("Error Connecting the database: ", err)
	// }
	// defer configs.CloseDatabase()

	// router.POST("/for-user", func(c *gin.Context) {
	// 	var data map[string]interface{}
	// 	err := c.BindJSON(&data)
	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	// Print the entire JSON data
	// 	fmt.Println("JSON data:", data)

	// 	// Access specific data using map keys (if known)
	// 	// for example:
	// 	// if name, ok := data["fields"].([]interface{})[0].(map[string]interface{})["name"]; ok {
	// 	// 	fmt.Println("Name:", name)
	// 	// }

	// 	// Loop over the map to access all key-value pairs
	// 	for key, value := range data {
	// 		fmt.Printf("Key: %s, Value: %v\n", key, value)
	// 	}
	// })

	// router.POST("/post-request", controllers.PostRequestHandler())

	// router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}
