package apierrors

import "net/http"

// MapErrorToHTTPStatus maps an ErrCode to an appropriate HTTP status code.
func MapErrorToHTTPStatus(err AppError) int {
    switch err.Code {
    case ErrCodeBadRequest:
        return http.StatusBadRequest // 400
    case ErrCodeUnauthorized:
        return http.StatusUnauthorized // 401
    case ErrCodeForbidden:
        return http.StatusForbidden // 403
    case ErrCodeNotFound:
        return http.StatusNotFound // 404
    case ErrCodeMethodNotAllowed:
        return http.StatusMethodNotAllowed // 405
    case ErrCodeConflict:
        return http.StatusConflict // 409
    case ErrCodeRequestTimeout:
        return http.StatusRequestTimeout // 408
    case ErrCodeGone:
        return http.StatusGone // 410
    case ErrCodeTooManyRequests:
        return http.StatusTooManyRequests // 429
    case ErrCodeInvalidContentType:
        return http.StatusUnsupportedMediaType // 415
    
    case ErrCodeValidationError, ErrCodeBindingError, ErrCodeInvalidInput:
        return http.StatusBadRequest // 400
    
    case ErrCodeDBError, ErrCodeTransactionFailure, ErrCodeDataConsistency:
        return http.StatusInternalServerError // 500
    case ErrCodeRecordNotFound:
        return http.StatusNotFound // 404
    case ErrCodeDuplicateRecord:
        return http.StatusConflict // 409
    case ErrCodeConstraintViolation:
        return http.StatusBadRequest // 400
    
    case ErrCodeServiceUnavailable, ErrCodeTimeout, ErrCodeDependencyFailure:
        return http.StatusServiceUnavailable // 503
    case ErrCodeNetworkError:
        return http.StatusBadGateway // 502
    case ErrCodeRateLimitExceeded:
        return http.StatusTooManyRequests // 429
    
    case ErrCodeFileNotFound:
        return http.StatusNotFound // 404
    case ErrCodeFileTooLarge:
        return http.StatusRequestEntityTooLarge // 413
    case ErrCodeFileUploadFailed, ErrCodeFileWriteError:
        return http.StatusInternalServerError // 500
    case ErrCodeUnsupportedFileType:
        return http.StatusUnsupportedMediaType // 415

    case ErrCodeAuthenticationFailed, ErrCodeTokenInvalid, ErrCodeTokenExpired, ErrCodeCSRFTokenInvalid:
        return http.StatusUnauthorized // 401
    case ErrCodeAuthorizationFailed:
        return http.StatusForbidden // 403

    case ErrCodeBusinessRule:
        return http.StatusUnprocessableEntity // 422

    case ErrCodeServerError:
        return http.StatusInternalServerError // 500
    case ErrCodeNotImplemented:
        return http.StatusNotImplemented // 501
    case ErrCodeBadGateway:
        return http.StatusBadGateway // 502
    case ErrCodeGatewayTimeout:
        return http.StatusGatewayTimeout // 504

    // Default case for custom or unknown errors.
    case ErrCodeCustom:
        return http.StatusInternalServerError // Default for custom errors
    case ErrCodeUnknown:
        return http.StatusInternalServerError // Catch-all for unknown errors

    default:
        return http.StatusInternalServerError // Default for unhandled cases
    }
}
