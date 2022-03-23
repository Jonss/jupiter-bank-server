package rest

func NewErrorResponses(errors ...ErrorResponse) ErrorResponses {
	var errResponses []ErrorResponse
	for _, v := range errors {
		errResponses = append(errResponses, v)
	}
	return ErrorResponses{errResponses}
}

const (
	CodeUser1          string = "USER-01" // user
	CodeUnexpected9999        = "J-9999"  // general
)

var UserExists = NewErrorResponses(ErrorResponse{Code: CodeUser1, Message: "user already exists"})
var UnexpectedError = NewErrorResponses(ErrorResponse{Code: CodeUnexpected9999, Message: "unexpected error occurred"})
