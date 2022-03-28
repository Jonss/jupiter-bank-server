package rest

import "strings"

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
