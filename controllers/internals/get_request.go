package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"restpad/restpad-service/configs"
	"restpad/restpad-service/utils"

	"github.com/gin-gonic/gin"
)

func GetRequestHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Make DB connection
		db, err := configs.OpenConnection()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": "Internal Server Error",
			})
			return
		}
		// Close the connection
		defer configs.CloseConnection()

		var query string
		condID := c.Param("cond_id")
		queryParams := c.Request.URL.Query()

		if condID != "" {
			query = utils.BuildGetQuery(condID)
		} else if c.Query("_page") != "" && c.Query("_limit") != "" {
			// Pagination block
			pageParam := c.Query("_page")
			limitParam := c.Query("_limit")

			page, err := strconv.Atoi(pageParam)
			if err != nil || page < 1 {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "error",
					"message": "Invalid page number",
				})
				return
			}

			limit, err := strconv.Atoi(limitParam)
			if err != nil || limit < 1 {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "error",
					"message": "Invalid limit value",
				})
				return
			}

			offset := (page - 1) * limit
			// Construct query with LIMIT and OFFSET for pagination
			query = fmt.Sprintf("SELECT * FROM demo ORDER BY id DESC LIMIT %d OFFSET %d;", limit, offset)
		} else if len(queryParams) > 0 {
			singleValueParams := utils.ConvertQueryParams(queryParams)
			query = utils.BuildGetQueryForFilters(singleValueParams)
		} else {
			query = "SELECT * FROM demo ORDER BY id DESC;"
		}

		// Execute the Query
		rows, err := db.Query(query)
		if rows == nil {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  "error",
				"message": "Record Not Found",
				"data":    utils.ResultRowsIfEmpty(nil),
			})
			return
		}
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Something went wrong",
			})
			return
		}
		defer rows.Close()

		resultRows, err := utils.ConvertRowsIntoValues(rows)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Something went wrong",
			})
			return
		}
		if (resultRows == nil && condID != "") || (resultRows == nil && len(queryParams) > 0) {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  "error",
				"message": "Record not found",
				"data":    utils.ResultRowsIfEmpty(resultRows),
			})
			return
		}
		fmt.Println(resultRows, "📌")

		// Return JSON response with array
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data":   utils.ResultRowsIfEmpty(resultRows),
		})
	}
}
