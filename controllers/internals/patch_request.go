package controllers

import (
	"net/http"
	"restpad/restpad-service/configs"
	"restpad/restpad-service/utils"

	"github.com/gin-gonic/gin"
)

func PatchRequestHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the condition ID from URL parameters and validate it.
		condID := c.Param("cond_id")
		if condID == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Missing ID parameter",
			})
			return
		}

		// Bind the JSON payload to a map.
		var data map[string]interface{}
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Invalid JSON payload: " + err.Error(),
			})
			return
		}

		// Generate the update query using a helper function.
		query := utils.BuildUpdateQuery(data, condID)
		// Optionally log the generated query here for debugging if needed.

		// Open a connection to the database.
		db, err := configs.OpenConnection()
		if err != nil {
			utils.HandleDBError(c, "Internal Server Error: unable to connect to database")
			return
		}
		// Ensure the connection is closed once the handler finishes.
		defer configs.CloseConnection()

		// Execute the query.
		res, err := db.Exec(query)
		if err != nil {
			utils.HandleDBError(c, "Database error: failed to execute update query")
			return
		}

		// Retrieve the number of affected rows.
		rowsAffected, err := res.RowsAffected()
		if err != nil {
			utils.HandleDBError(c, "Database error: unable to retrieve affected rows")
			return
		}
		if rowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"status":        "error",
				"message":       "No record found",
				"rows_affected": rowsAffected,
			})
			return
		}

		// Return a success response.
		c.JSON(http.StatusOK, gin.H{
			"status":        "success",
			"message":       "Record updated successfully",
			"rows_affected": rowsAffected,
		})
	}
}
