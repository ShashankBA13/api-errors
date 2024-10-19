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
    return APIErrorResponse{
        StatusCode: statusCode,
        Message:    message,
        AppErrors:  errs,
    }
}
