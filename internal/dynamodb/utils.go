package dynamodb

import (
	"encoding/json"
)

// JSONStringify - returns a json string representation of the given value
func JSONStringify(v any) string {
	data, err := json.Marshal(v)
	if err != nil {
		return ""
	}

	return string(data)
}
