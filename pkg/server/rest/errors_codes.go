package rest

import "strings"

const (
	CodeValidation1        = "VAL-0001"  // general
	CodeUser1          string = "USER-0001" // user
	CodeUnexpected9999        = "J-9999"  // general
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
