package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseWithJson(writer http.ResponseWriter, status int, object interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	//writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	writer.WriteHeader(status)
	if err := json.NewEncoder(writer).Encode(object); err != nil {
		return
	}
}
