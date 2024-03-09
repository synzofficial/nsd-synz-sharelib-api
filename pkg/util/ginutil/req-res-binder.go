package ginutil

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	defaultmodel "github.com/synzofficial/nsd-synz-sharelib-api/pkg/model/default-model"
	errormodel "github.com/synzofficial/nsd-synz-sharelib-api/pkg/model/error-model"
)

// 200 OK:
// The request succeeded. The result meaning of "success" depends on the HTTP method:
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/200
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

	c.JSON(http.StatusOK, defaultmodel.RaiseSuccessResponse(http.StatusOK, resp))
}

// 201 Created:
// The request succeeded, and a new resource was created as a result. This is typically the response sent after POST requests, or some PUT requests.
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/201
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

	c.JSON(http.StatusCreated, defaultmodel.RaiseSuccessResponse(http.StatusCreated, resp))
}

// 204 No Content:
// There is no content to send for this request, but the headers may be useful. The user agent may update its cached headers for this resource with the new ones.
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/204
func BindReqJson204Resp[Request any](c *gin.Context, f func(ctx context.Context, request Request) error) {
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

// 200 ok but no content
// use for some service need to return 200 status but no response body.
func BindReqJson200RespNoContent[Request any](c *gin.Context, f func(ctx context.Context, request Request) error) {
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

	c.JSON(http.StatusOK, defaultmodel.RaiseSuccessResponse(http.StatusOK, nil))
}
