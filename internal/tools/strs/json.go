package strs

import "encoding/json"

// ToJson converts a given object to its JSON string representation.
// Note: This function ignores any errors that occur during the marshaling process.
func ToJson(obj interface{}) string {
	jsonData, _ := json.Marshal(obj)
	return string(jsonData)
}
