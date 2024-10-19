package apierrors

// ErrCode represents error codes for various error scenarios.
type ErrCode string

// HTTP-specific error codes.
const (
	// Client-side errors (400 range).
	ErrCodeBadRequest         ErrCode = "BAD_REQUEST"          // 400 - General client-side error.
	ErrCodeUnauthorized       ErrCode = "UNAUTHORIZED"         // 401 - Authentication is required and has failed or has not yet been provided.
	ErrCodeForbidden          ErrCode = "FORBIDDEN"            // 403 - Client does not have permission to access this resource.
	ErrCodeNotFound           ErrCode = "NOT_FOUND"            // 404 - Requested resource could not be found.
	ErrCodeMethodNotAllowed   ErrCode = "METHOD_NOT_ALLOWED"   // 405 - HTTP method not supported.
	ErrCodeRequestTimeout     ErrCode = "REQUEST_TIMEOUT"      // 408 - Request took too long.
	ErrCodeConflict           ErrCode = "CONFLICT"             // 409 - Resource conflict, like duplicate data.
	ErrCodeGone               ErrCode = "GONE"                 // 410 - Resource is no longer available.
	ErrCodeTooManyRequests    ErrCode = "TOO_MANY_REQUESTS"    // 429 - Rate limiting error.
	ErrCodeInvalidContentType ErrCode = "INVALID_CONTENT_TYPE" // Unsupported content type in request.

	// Server-side errors (500 range).
	ErrCodeServerError        ErrCode = "SERVER_ERROR"        // 500 - Generic server error.
	ErrCodeNotImplemented     ErrCode = "NOT_IMPLEMENTED"     // 501 - Functionality not implemented.
	ErrCodeBadGateway         ErrCode = "BAD_GATEWAY"         // 502 - Received an invalid response from upstream server.
	ErrCodeServiceUnavailable ErrCode = "SERVICE_UNAVAILABLE" // 503 - Service is temporarily unavailable.
	ErrCodeGatewayTimeout     ErrCode = "GATEWAY_TIMEOUT"     // 504 - Timeout in gateway or proxy.
)

// Application-specific error codes.
const (
	ErrCodeValidationError    ErrCode = "VALIDATION_ERROR"       // Input validation failed.
	ErrCodeBindingError       ErrCode = "BINDING_ERROR"          // Error binding request parameters.
	ErrCodeInvalidInput       ErrCode = "INVALID_INPUT"          // Input does not match the required format.
	ErrCodeResourceExhausted  ErrCode = "RESOURCE_EXHAUSTED"     // Resource limits have been exceeded, e.g., storage quota.
	ErrCodeBusinessRule       ErrCode = "BUSINESS_RULE_ERROR"    // Business logic condition not met.
	ErrCodeDeprecationWarning ErrCode = "DEPRECATION_WARNING"    // Using deprecated features or API versions.
	ErrCodeDataConsistency    ErrCode = "DATA_CONSISTENCY_ERROR" // Data integrity issues like inconsistent data states.
)

// Database-related error codes.
const (
	ErrCodeDBError             ErrCode = "DB_ERROR"             // General database error.
	ErrCodeRecordNotFound      ErrCode = "RECORD_NOT_FOUND"     // Specific database record could not be found.
	ErrCodeDuplicateRecord     ErrCode = "DUPLICATE_RECORD"     // Duplicate record insertion.
	ErrCodeTransactionFailure  ErrCode = "TRANSACTION_FAILURE"  // Database transaction failed.
	ErrCodeConstraintViolation ErrCode = "CONSTRAINT_VIOLATION" // Constraint check failed (e.g., foreign key).
)

// Integration and API-to-API communication error codes.
const (
	ErrCodeTimeout           ErrCode = "TIMEOUT_ERROR"       // Timeout during API call to external service.
	ErrCodeRateLimitExceeded ErrCode = "RATE_LIMIT_EXCEEDED" // Rate-limiting error when interacting with external services.
	ErrCodeNetworkError      ErrCode = "NETWORK_ERROR"       // General network error (e.g., DNS failure).
	ErrCodeDependencyFailure ErrCode = "DEPENDENCY_FAILURE"  // Failed due to external service dependency.
	ErrCodeInvalidResponse   ErrCode = "INVALID_RESPONSE"    // Received invalid or unexpected response from an external service.
)

// Security-related error codes.
const (
	ErrCodeAuthenticationFailed ErrCode = "AUTHENTICATION_FAILED" // Failed authentication (e.g., incorrect password).
	ErrCodeAuthorizationFailed  ErrCode = "AUTHORIZATION_FAILED"  // User lacks required permissions.
	ErrCodeTokenExpired         ErrCode = "TOKEN_EXPIRED"         // Authentication token has expired.
	ErrCodeTokenInvalid         ErrCode = "TOKEN_INVALID"         // Provided token is invalid.
	ErrCodeCSRFTokenInvalid     ErrCode = "CSRF_TOKEN_INVALID"    // Invalid CSRF token for request.
)

// File-related error codes.
const (
	ErrCodeFileNotFound        ErrCode = "FILE_NOT_FOUND"        // Requested file could not be found.
	ErrCodeFileUploadFailed    ErrCode = "FILE_UPLOAD_FAILED"    // Error during file upload.
	ErrCodeFileTooLarge        ErrCode = "FILE_TOO_LARGE"        // Uploaded file size exceeds the allowed limit.
	ErrCodeUnsupportedFileType ErrCode = "UNSUPPORTED_FILE_TYPE" // File type is not supported.
	ErrCodeFileReadError       ErrCode = "FILE_READ_ERROR"       // Error reading file content.
	ErrCodeFileWriteError      ErrCode = "FILE_WRITE_ERROR"      // Error writing to file.
)

// Custom or unknown error codes.
const (
	ErrCodeCustom  ErrCode = "CUSTOM_ERROR"  // For custom or specific use-case errors.
	ErrCodeUnknown ErrCode = "UNKNOWN_ERROR" // Catch-all for unclassified errors.
)
