package utils

import (
	"fmt"
	"strings"
)

// util function to return if string the in ” or 5
func formatValue(value interface{}) string {
	switch v := value.(type) {
	case string:
		return fmt.Sprintf("'%s'", v)
	default:
		return fmt.Sprintf("%v", v)
	}
}

// Function to build Get Query
func BuildGetQuery(condID string) string {
	final := fmt.Sprintf("SELECT * FROM demo WHERE id = %s;", condID)

	return final
}

// Function to build Insert Query
func BuildInsertQuery(data map[string]interface{}) string {
	var columns []string
	var values []string

	for column, value := range data {
		columns = append(columns, column)
		values = append(values, formatValue(value))
	}

	columnsStr := strings.Join(columns, ", ")
	valuesStr := strings.Join(values, ", ")

	final := fmt.Sprintf("INSERT INTO demo (%s) VALUES (%s);", columnsStr, valuesStr)
	return final
}

// Function to build Update Query
func BuildUpdateQuery(data map[string]interface{}, condID string) string {
	var setClauses []string

	for column, value := range data {
		if column != "id" {
			setClauses = append(setClauses, fmt.Sprintf("%s = %s", column, formatValue(value)))
		}
	}

	setStr := strings.Join(setClauses, ", ")
	final := fmt.Sprintf("UPDATE demo SET %s WHERE id = %s;", setStr, condID)

	return final
}

// Function to build Update Query
func BuildDeleteQuery(condID string) string {
	final := fmt.Sprintf("DELETE FROM demo WHERE id = %s;", condID)

	return final
}

// Function to build Put Query
func BuildPutQuery(data map[string]interface{}, condID string) string {
	final := BuildGetQuery(condID)

	return final
}
