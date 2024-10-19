package apierrors

import (
	"fmt"
	"log"
)

type AppError struct {
	Stack   *stackTrace `json:"-"`
	Message string      `json:"message"`
	Code    ErrCode     `json:"code"`
}

func (e AppError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// NewAppError creates a new instance of AppError.
func NewAppError(code ErrCode, message string) AppError {
	return AppError{
		Stack:   collectStackTrace(),
		Message: message,
		Code:    code,
	}
}

// LogError logs the error's stack trace for server-side debugging.
func (e AppError) LogError() {
	if e.Stack != nil {
		log.Printf("Error: %s, Code: %s\nStack Trace:\n%s", e.Message, e.Code, e.Stack)
	} else {
		log.Printf("Error: %s, Code: %s", e.Message, e.Code)
	}
}
