package errorutil

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/synzofficial/nsd-synz-sharelib-api/pkg/model/errormodel"
)

func CastErrorJson(c *gin.Context, err error) {
	switch e := err.(type) {
	case *errormodel.CustomError:
		c.JSON(e.Status, e)
	default:
		c.JSON(http.StatusInternalServerError, e)
	}
}
