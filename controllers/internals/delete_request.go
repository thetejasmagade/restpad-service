package controllers

import (
	"fmt"
	"net/http"
	"restpad/restpad-service/configs"
	"restpad/restpad-service/utils"

	"github.com/gin-gonic/gin"
)

func DeleteRequestHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data map[string]interface{}

		fmt.Println(c.Param("id"))
		fmt.Println(c.Request.Header["Authorization"])

		// Bind the JSON
		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// GENERATE Delete query
		query := utils.BuildDeleteQuery(c.Param("cond_id"))

		// Make a connection to DB
		db, err := configs.OpenConnection()
		if err != nil {
			utils.HandleDBError(c, "Internal Server Error")
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
		c.JSON(http.StatusOK, gin.H{
			"status":        http.StatusOK,
			"message":       "Record Deleted Successfully.",
			"rows_affected": rowsAffectedCnt,
		})
	}
}
