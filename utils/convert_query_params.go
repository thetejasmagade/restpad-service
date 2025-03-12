package utils

import "net/url"

// This function helps BuildGetQueryForFilters function. convert map[string][]string to map[string]string
func ConvertQueryParams(queryParams url.Values) map[string]string {
	singleValueParams := make(map[string]string)
	for key, values := range queryParams {
		if len(values) > 0 {
			singleValueParams[key] = values[0] // Take the first value only
		}
	}
	return singleValueParams
}
