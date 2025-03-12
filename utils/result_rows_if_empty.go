package utils

func ResultRowsIfEmpty(resultRows []map[string]interface{}) []map[string]interface{} {
	if resultRows == nil {
		return []map[string]interface{}{} // Return an empty array instead of nil
	}
	return resultRows
}
