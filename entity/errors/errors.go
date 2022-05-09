package errors

import "fmt"

// ResourceNotFoundError ...
type ResourceNotFoundError struct {
	Resource string
}

// Error ...
func (e ResourceNotFoundError) Error() string {
	return fmt.Sprintf("%s not found", e.Resource)
}

// NewResourceNotFoundError ...
func NewResourceNotFoundError(message string) error {
	return ResourceNotFoundError{message}
}

// BusinessError ...
type BusinessError struct {
	Message string
}

// Error ...
func (e BusinessError) Error() string {
	return e.Message
}

// NewBusinessError ...
func NewBusinessError(message string) error {
	return BusinessError{message}
}
