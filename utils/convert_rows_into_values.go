package utils

import (
	"database/sql"
)

func ConvertRowsIntoValues(rows *sql.Rows) ([]map[string]interface{}, error) {
	// Get the columns
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	// Declares a variable resultRows as a slice of maps, where each map represents a row from the database with column names as keys and corresponding values.
	var resultRows []map[string]interface{}

	// Creates a slice named values of empty interfaces (interface{}) with a length equal to the number of columns in the result set.
	values := make([]interface{}, len(columns))
	// Creates another slice named valuePtrs of empty interfaces with the same length as values. This slice will be used to store pointers to the elements in the values slice.
	valuePtrs := make([]interface{}, len(columns))

	// Iterates over the columns and assigns the address of each element in the values slice to the corresponding element in the valuePtrs slice. This prepares the pointers for the rows.Scan method.
	for i := range columns {
		valuePtrs[i] = &values[i]
	}

	// Starts a loop to iterate over the result set rows using rows.Next().
	for rows.Next() {
		// Scans the current row of the result set and assigns the values to the variables pointed to by the pointers in the valuePtrs slice. The ... syntax is used to pass a slice as individual arguments to the Scan method.
		err := rows.Scan(valuePtrs...)
		if err != nil {
			return nil, err
		}

		// Initializes an empty map named rowData to store the values of the current row.
		rowData := make(map[string]interface{})

		// Iterates over the columns and assigns each column value to the corresponding key in the rowData map.
		for i, col := range columns {
			rowData[col] = values[i]
		}

		// Appends the rowData map (representing the current row) to the resultRows slice.
		resultRows = append(resultRows, rowData)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return resultRows, nil
}
