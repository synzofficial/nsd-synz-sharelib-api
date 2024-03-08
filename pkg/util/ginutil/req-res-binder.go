package ginutil

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/synzofficial/nsd-synz-sharelib-api/pkg/model/errormodel"
)

func BindReqJson200Resp[Request any, Response any](c *gin.Context, f func(ctx context.Context, request Request) (Response, error)) {
	var request Request

	ctx, request, err := BindingRequestLogging(c, request)
	if err != nil {
		CastErrorJsonWithLogging(c, errormodel.RaiseBindingBadRequestError(err))
		return
	}

	resp, err := f(ctx, request)
	if err != nil {
		CastErrorJsonWithLogging(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func BindReqJson201Resp[Request any, Response any](c *gin.Context, f func(ctx context.Context, request Request) (Response, error)) {
	var request Request

	ctx, request, err := BindingRequestLogging(c, request)
	if err != nil {
		CastErrorJsonWithLogging(c, errormodel.RaiseBindingBadRequestError(err))
		return
	}

	resp, err := f(ctx, request)
	if err != nil {
		CastErrorJsonWithLogging(c, err)
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func BindReqJson204Resp[Request any, Response any](c *gin.Context, f func(ctx context.Context, request Request) error) {
	var request Request

	ctx, request, err := BindingRequestLogging(c, request)
	if err != nil {
		CastErrorJsonWithLogging(c, errormodel.RaiseBindingBadRequestError(err))
		return
	}

	if err := f(ctx, request); err != nil {
		CastErrorJsonWithLogging(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func BindReqJson200RespNoContent[Request any, Response any](c *gin.Context, f func(ctx context.Context, request Request) error) {
	var request Request

	ctx, request, err := BindingRequestLogging(c, request)
	if err != nil {
		CastErrorJsonWithLogging(c, errormodel.RaiseBindingBadRequestError(err))
		return
	}

	if err := f(ctx, request); err != nil {
		CastErrorJsonWithLogging(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
