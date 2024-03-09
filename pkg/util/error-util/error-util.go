package errorutil

import (
	"net/http"

	"github.com/gin-gonic/gin"
	errormodel "github.com/synzofficial/nsd-synz-sharelib-api/pkg/model/error-model"
)

func CastErrorJson(c *gin.Context, err error) {
	switch e := err.(type) {
	case *errormodel.CustomError:
		c.JSON(e.Status, e)
	default:
		c.JSON(http.StatusInternalServerError, errormodel.CustomError{
			Status:  http.StatusInternalServerError,
			Code:    "9999",
			Message: e.Error(),
		})
	}
}
