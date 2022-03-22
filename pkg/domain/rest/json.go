package rest

import (
	"encoding/json"
	"net/http"
)

const contentType = "Content-Type"
const applicationJson = "application/json"

func JsonResponse(w http.ResponseWriter, statusCode int, responseBody interface{}) {
	response, _ := json.Marshal(responseBody)

	w.Header().Set(contentType, applicationJson)
	w.WriteHeader(statusCode)
	w.Write(response)
}
