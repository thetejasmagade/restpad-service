package controllers

import (
	"net/http"
	"restpad/restpad-service/configs"
	"restpad/restpad-service/utils"

	"github.com/gin-gonic/gin"
)

func PutRequestHandler() gin.HandlerFunc {
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
		updateQuery := utils.BuildUpdateQuery(data, condID)

		// Open a connection to the database.
		db, err := configs.OpenConnection()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": "Internal Server Error",
			})
			return
		}
		// Ensure the connection is closed once the handler finishes.
		defer configs.CloseConnection()

		// Execute the update query.
		res, err := db.Exec(updateQuery)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Something went wrong",
			})
			return
		}

		// Retrieve the number of affected rows.
		rowsAffected, err := res.RowsAffected()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Something went wrong",
			})
			return
		}

		// If no record was updated, then insert a new record.
		if rowsAffected == 0 {
			insertQuery := utils.BuildInsertQuery(data)
			resInsert, err := db.Exec(insertQuery)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "error",
					"message": "Something went wrong",
				})
				return
			}
			rowsInserted, err := resInsert.RowsAffected()
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "error",
					"message": "Something went wrong",
				})
				return
			}
			if rowsInserted == 0 {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":        "error",
					"message":       "Failed to insert new record",
					"rows_affected": rowsInserted,
				})
				return
			}
			c.JSON(http.StatusCreated, gin.H{
				"status":        "success",
				"message":       "New record inserted successfully",
				"rows_affected": rowsInserted,
			})
			return
		}

		// Return a success response if the update was successful.
		c.JSON(http.StatusOK, gin.H{
			"status":        "success",
			"message":       "Record updated successfully",
			"rows_affected": rowsAffected,
		})
	}
}
