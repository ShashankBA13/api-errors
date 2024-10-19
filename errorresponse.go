package apierrors

import "encoding/json"

type APIErrorResponse struct {
	StatusCode int        `json:"status_code"`
	Message    string     `json:"message"`
	AppErrors  []AppError `json:"errors"`
}

func (r *APIErrorResponse) AddError(err AppError) {
	r.AppErrors = append(r.AppErrors, err)
}

func (r *APIErrorResponse) ToJSON() ([]byte, error) {
	return json.Marshal(r)
}

func NewAPIErrorResponse(statusCode int, message string, errs ...AppError) APIErrorResponse {
	// Exclude the stack trace from the errors before returning them to the client.
	strippedErrors := make([]AppError, len(errs))
	for i, err := range errs {
		strippedErrors[i] = AppError{
			Message: err.Message,
			Code:    err.Code,
		}
	}
	return APIErrorResponse{
		StatusCode: statusCode,
		Message:    message,
		AppErrors:  strippedErrors,
	}
}
