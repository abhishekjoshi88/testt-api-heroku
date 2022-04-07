package domain

// ErrorDetails is a struct used for storing response of error details
type ErrorDetails struct {
	Code        string `json:"errorCode"`
	Description string `json:"errorDescription"`
}

var UnexpectedError = ErrorDetails{Code: "unexpectedError", Description: "An unexpected error occurred. Please try again later."}
