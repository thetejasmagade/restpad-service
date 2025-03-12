package controllers

import (
	"fmt"
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

		var query string
		condID := c.Param("cond_id")
		queryParams := c.Request.URL.Query()
		fmt.Println(queryParams)

		if condID != "" {
			query = utils.BuildGetQuery(condID)
		} else if c.Query("_page") != "" && c.Query("_limit") != "" {

		} else if len(queryParams) > 0 {
			singleValueParams := utils.ConvertQueryParams(queryParams)
			query = utils.BuildGetQueryForFilters(singleValueParams)
		} else {
			query = "SELECT * FROM demo ORDER BY id DESC;"
		}
		// Execute the Query
		rows, err := db.Query(query)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  "error",
				"message": "Data not found",
			})
			return
		}
		defer rows.Close()

		resultRows, err := utils.ConvertRowsIntoValues(rows)
		if err != nil {
			utils.HandleDBError(c, "Internal Server Error")
			return
		}
		if resultRows == nil && condID != "" {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  "error",
				"message": "Data not found",
				"data":    utils.ResultRowsIfEmpty(resultRows),
				// "rows_affected": len(resultRows), // Number of selected rows
			})
			return
		}
		fmt.Println(resultRows, "📌")

		// fmt.Printf("Rows selected: %d\n", len(resultRows))
		// fmt.Println("Selected data:", resultRows)

		// Return JSON response with array
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data":   utils.ResultRowsIfEmpty(resultRows),
			// "rows_affected": len(resultRows), // Number of selected rows
		})
	}
}
