package util

import "encoding/json"

func MustJSON(obj any) string {
	bytes, _ := json.Marshal(obj)
	return string(bytes)
}

func MustJSONIndented(obj any) string {
	//bytes, _ := json.MarshalIndent(obj, "", " ")
	bytes, _ := json.Marshal(obj)
	return string(bytes)
}
