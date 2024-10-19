package apierrors

import "fmt"

type AppError struct {
	Stack   *stackTrace `json:"stack,omitempty"`
	Message string      `json:"message"`
	Code    ErrCode     `json:"code"`
	Field   string      `json:"field,omitempty"`
	Value   string      `json:"value,omitempty"`
}

func (e AppError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// NewAppError creates a new instance of AppError.
func NewAppError(code ErrCode, message, field, value string) AppError {
	return AppError{
		Stack:   collectStackTrace(),
		Message: message,
		Code:    code,
		Field:   field,
		Value:   value,
	}
}
