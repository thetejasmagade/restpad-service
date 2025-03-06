package controllers

import (
	"net/http"
	"restpad/restpad-service/configs"
	"restpad/restpad-service/utils"

	"github.com/gin-gonic/gin"
)

func GetRequestHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Make DB connection
		db, err := configs.OpenConnection()
		if err != nil {
			utils.HandleDBError(c, "Internal Server Error")
			// handleDBError(c, "Internal Server Error")
			return
		}
		// Close the connection
		defer configs.CloseConnection()

		// Execute the Query
		rows, err := db.Query("SELECT * FROM demo ORDER BY id ASC;")
		if err != nil {
			utils.HandleDBError(c, "Internal Server Error")
			return
		}
		defer rows.Close()

		resultRows, err := utils.ConvertRowsIntoValues(rows)
		if err != nil {
			utils.HandleDBError(c, "Internal Server Error")
			return
		}

		// fmt.Printf("Rows selected: %d\n", len(resultRows))
		// fmt.Println("Selected data:", resultRows)

		// Return JSON response with array
		c.JSON(http.StatusOK, gin.H{
			"status":        http.StatusOK,
			"data":          resultRows,
			"rows_affected": len(resultRows), // Number of selected rows
		})
	}
}
