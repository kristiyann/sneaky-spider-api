package models

import (
	"net/http"
	"strconv"
)

type APIError struct {
	Message    string
	HttpStatus int
}

type APIErrorResponse struct {
	Errors     []string `json:"errors"`
	HttpStatus int      `json:"http_status"`
}

type APIErrorWithTraceResponse struct {
	APIErrorResponse
	TraceID string `json:"trace_id"`
}

func (e *APIError) Error() string {
	return e.Message
}

func NewAPIError(message string, status int) *APIError {
	return &APIError{
		Message:    strconv.Itoa(status) + ": " + message,
		HttpStatus: status,
	}
}

func NewInvalidParameters() *APIError {
	return NewAPIError("invalid parameters", http.StatusBadRequest)
}
