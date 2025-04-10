package controllers

import (
	"fmt"
	"net/http"
	"restpad/restpad-service/configs"
	"restpad/restpad-service/utils"

	"github.com/gin-gonic/gin"
)

func PostRequestHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data map[string]interface{}

		fmt.Println(c.Param("id"))
		fmt.Println(c.Request.Header["Authorization"])

		// Bind the JSON
		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Invalid JSON payload: " + err.Error(),
			})
			return
		}

		// GENERATE insert query
		query := utils.BuildInsertQuery(data, "demo")

		// Make a connection to DB
		db, err := configs.OpenConnection()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": "Internal Server Error",
			})
			return
		}
		defer configs.CloseConnection()

		// Execute the Query
		res, err := db.Exec(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Something went wrong",
			})
			return
		}

		// get RowsAffected
		rowsAffectedCnt, err := res.RowsAffected()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Something went wrong",
			})
			return
		}

		// return response or error if any
		// message := fmt.Sprintf("%d rows affected", rowsAffectedCnt)
		c.JSON(http.StatusCreated, gin.H{
			"status":        http.StatusCreated,
			"message":       "Record Inserted Successfully.",
			"rows_affected": rowsAffectedCnt,
		})
	}
}
