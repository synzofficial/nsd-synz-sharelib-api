package basemodel

import (
	"github.com/synzofficial/nsd-synz-sharelib-api/pkg/enum"
)

const (
	code    = "0"
	message = "success"
)

type SuccessResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func RaiseSuccessResponse(data any) SuccessResponse {

	if d, ok := (data).(SuccessResponse); ok {
		return d
	}

	if d, ok := (data).(*SuccessResponse); ok {
		return *d
	}

	return SuccessResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func RaiseSuccessCreateResponse() *SuccessResponse {
	return &SuccessResponse{
		Code:    code,
		Message: message,
		Data:    enum.SuccessType_CREATED,
	}
}

func RaiseSuccessUpdateResponse() *SuccessResponse {
	return &SuccessResponse{
		Code:    code,
		Message: message,
		Data:    enum.SuccessType_UPDATED,
	}
}

func RaiseSuccessDeleteResponse() *SuccessResponse {
	return &SuccessResponse{
		Code:    code,
		Message: message,
		Data:    enum.SuccessType_DELETED,
	}
}
