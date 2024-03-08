package errormodel

import (
	"fmt"
	"net/http"
)

// base error code
const (
	UnauthorizedCode = "0001"
	BadRequestCode   = "0002"
)

type CustomError struct {
	Status  int    `json:"-"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func UnauthorizedError(err error) *CustomError {
	return &CustomError{
		Status:  http.StatusUnauthorized,
		Code:    UnauthorizedCode,
		Message: err.Error(),
	}
}

func RaiseBindingBadRequestError(err error) *CustomError {
	return &CustomError{
		Status:  http.StatusBadRequest,
		Code:    BadRequestCode,
		Message: err.Error(),
	}
}

func RaiseBadRequestError(code string, msg string) *CustomError {
	return &CustomError{
		Status:  http.StatusBadRequest,
		Code:    code,
		Message: msg,
	}
}

func RaiseBusinessError(code string, msg string) *CustomError {
	return &CustomError{
		Status:  http.StatusConflict,
		Code:    code,
		Message: msg,
	}
}

func RaiseTechnicalError(code string, msg string) *CustomError {
	return &CustomError{
		Status:  http.StatusInternalServerError,
		Code:    code,
		Message: msg,
	}
}

func RaiseCustomError(status int, code string, msg string) *CustomError {
	return &CustomError{
		Status:  status,
		Code:    code,
		Message: msg,
	}
}
