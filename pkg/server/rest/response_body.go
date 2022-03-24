package rest

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"net/http"
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
	JsonResponse(w, http.StatusBadRequest, NewErrorResponses(errorResponses...))
}

