package apierrors

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// HandleDBError processes different types of database errors and returns appropriate responses.
func HandleDBError(ctx *gin.Context, err error) {
	var apperror AppError

	// Handle specific database error types
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		apperror = NewAppError(ErrCodeRecordNotFound, "Record not found")
		apperror.LogError() // Log the error with stack trace on the server
		apiErrorRsp := NewAPIErrorResponse(http.StatusNotFound, "The requested resource could not be found", apperror)
		ctx.JSON(apiErrorRsp.StatusCode, apiErrorRsp)

	case errors.Is(err, gorm.ErrInvalidData):
		apperror = NewAppError(ErrCodeBadRequest, "Invalid data")
		apperror.LogError()
		apiErrorRsp := NewAPIErrorResponse(http.StatusBadRequest, "Invalid data provided to the database", apperror)
		ctx.JSON(apiErrorRsp.StatusCode, apiErrorRsp)

	case errors.Is(err, gorm.ErrInvalidTransaction):
		apperror = NewAppError(ErrCodeTransactionFailure, "Invalid transaction")
		apperror.LogError()
		apiErrorRsp := NewAPIErrorResponse(http.StatusInternalServerError, "There was an issue with the transaction", apperror)
		ctx.JSON(apiErrorRsp.StatusCode, apiErrorRsp)

	case errors.Is(err, gorm.ErrInvalidDB):
		apperror = NewAppError(ErrCodeServiceUnavailable, "Database connection issue")
		apperror.LogError()
		apiErrorRsp := NewAPIErrorResponse(http.StatusServiceUnavailable, "Database connection is currently unavailable", apperror)
		ctx.JSON(apiErrorRsp.StatusCode, apiErrorRsp)

	default:
		apperror = NewAppError(ErrCodeDBError, "Database error")
		apperror.LogError()
		apiErrorRsp := NewAPIErrorResponse(http.StatusInternalServerError, "Something went wrong with the database", apperror)
		ctx.JSON(apiErrorRsp.StatusCode, apiErrorRsp)
	}
}
