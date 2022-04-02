package server

import (
	"encoding/json"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ErrorResponses struct {
	Errors []ErrorResponse `json:"errors"`
}

func ValidateRequestBody(err error, w http.ResponseWriter, translator ut.Translator) {
	validationErrors := err.(validator.ValidationErrors)
	errorResponses := make([]ErrorResponse, len(validationErrors))
	for i, vErr := range validationErrors {
		errorResponses[i] = NewValidationError(vErr.Translate(translator))
	}
	apiResponse(w, http.StatusBadRequest, NewErrorResponses(errorResponses...))
}

const contentType = "Content-Type"
const applicationJson = "application/json"

func apiResponse(w http.ResponseWriter, statusCode int, responseBody interface{}) {
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

// errors
const (
	CodeValidation1           = "VAL-0001"  // validation
	CodeAuth1                 = "AUTH-0001" // Authentication
	CodeUser1          string = "USER-0001" // user
	CodeUser2          string = "USER-0002" // user unauthorized
	CodeUnexpected9999        = "J-9999"    // general
)

func NewErrorResponses(errors ...ErrorResponse) ErrorResponses {
	var errResponses []ErrorResponse
	for _, v := range errors {
		errResponses = append(errResponses, v)
	}
	return ErrorResponses{errResponses}
}

func NewValidationError(message string) ErrorResponse {
	return ErrorResponse{Code: CodeValidation1, Message: strings.ToLower(message)}
}

var UserExists = NewErrorResponses(ErrorResponse{Code: CodeUser1, Message: "user already exists"})
var UnexpectedError = NewErrorResponses(ErrorResponse{Code: CodeUnexpected9999, Message: "unexpected error occurred"})
var Unauthorized = NewErrorResponses(ErrorResponse{Code: CodeUser2, Message: "unauthorized"})
var AuthorizationIsInvalid = NewErrorResponses(ErrorResponse{Code: CodeUser2, Message: "Authorization header is invalid"})
