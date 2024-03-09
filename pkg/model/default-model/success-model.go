package defaultmodel

import (
	"github.com/synzofficial/nsd-synz-sharelib-api/pkg/enum"
)

const (
	code    = "0"
	message = "success"
)

type SuccessResponse struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func RaiseSuccessResponse(status int, data any) SuccessResponse {

	if d, ok := (data).(SuccessResponse); ok {
		return d
	}

	if d, ok := (data).(*SuccessResponse); ok {
		return *d
	}

	return SuccessResponse{
		Status:  status,
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func RaiseSuccessCreateResponse(status int) *SuccessResponse {
	return &SuccessResponse{
		Status:  status,
		Code:    code,
		Message: message,
		Data:    enum.SuccessType_CREATED,
	}
}

func RaiseSuccessUpdateResponse(status int) *SuccessResponse {
	return &SuccessResponse{
		Status:  status,
		Code:    code,
		Message: message,
		Data:    enum.SuccessType_UPDATED,
	}
}

func RaiseSuccessDeleteResponse(status int) *SuccessResponse {
	return &SuccessResponse{
		Status:  status,
		Code:    code,
		Message: message,
		Data:    enum.SuccessType_DELETED,
	}
}
