package rest

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ErrorResponses struct {
	Errors []ErrorResponse `json:"errors"`
}
