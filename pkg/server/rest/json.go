package rest

import (
	"encoding/json"
	"net/http"
)

const contentType = "Content-Type"
const applicationJson = "application/json"

func JsonResponse(w http.ResponseWriter, statusCode int, responseBody interface{}) {
	response, err := json.Marshal(responseBody)
	if err != nil {
		panic(err)
	}

	w.Header().Set(contentType, applicationJson)
	w.WriteHeader(statusCode)
	_, err = w.Write(response)
	if err != nil {
		panic(err)
	}
}
