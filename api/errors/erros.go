package errors

import (
	"fmt"
	"net/http"
)

// ErrorResponse ...
type ErrorResponse struct {
	Status  int         `json:"status_code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

// Error ...
func (e ErrorResponse) Error() string {
	return e.Message
}

// StatusCode ...
func (e ErrorResponse) StatusCode() int {
	return e.Status
}

// InternalServerError ...
func InternalServerError(msg interface{}) ErrorResponse {
	if msg == "" {
		msg = "A system or application error occurred."
	}
	return ErrorResponse{
		Status:  http.StatusInternalServerError,
		Message: fmt.Sprintf("%v", msg),
	}
}

// NotFound ...
func NotFound(msg interface{}) ErrorResponse {
	if msg == "" {
		msg = "The specified resource does not found."
	}
	return ErrorResponse{
		Status:  http.StatusNotFound,
		Message: fmt.Sprintf("%v", msg),
	}
}

// Unauthorized ...
func Unauthorized(msg interface{}) ErrorResponse {
	if msg == "" {
		msg = "Client authentication failed."
	}
	return ErrorResponse{
		Status:  http.StatusUnauthorized,
		Message: fmt.Sprintf("%v", msg),
	}
}

// Forbidden ...
func Forbidden(msg interface{}) ErrorResponse {
	if msg == "" {
		msg = "Authorization failed due to insufficient permissions."
	}
	return ErrorResponse{
		Status:  http.StatusForbidden,
		Message: fmt.Sprintf("%v", msg),
	}
}

// BadRequest ...
func BadRequest(msg interface{}) ErrorResponse {
	if msg == "" {
		msg = "Request is not well-formed, syntactically incorrect, or violates schema."
	}
	return ErrorResponse{
		Status:  http.StatusBadRequest,
		Message: fmt.Sprintf("%v", msg),
	}
}

// UnprocessableEntity ...
func UnprocessableEntity(msg interface{}) ErrorResponse {
	if msg == "" {
		msg = "The request is semantically incorrect or fails business validation."
	}
	return ErrorResponse{
		Status:  http.StatusUnprocessableEntity,
		Message: fmt.Sprintf("%v", msg),
	}
}
