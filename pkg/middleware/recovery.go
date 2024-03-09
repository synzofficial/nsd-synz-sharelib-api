package middleware

import (
	"bytes"

	"github.com/gin-gonic/gin"
	errormodel "github.com/synzofficial/nsd-synz-sharelib-api/pkg/model/error-model"
	errorutil "github.com/synzofficial/nsd-synz-sharelib-api/pkg/util/error-util"
)

func Recovery() gin.HandlerFunc {
	errLog := new(bytes.Buffer)
	return gin.RecoveryWithWriter(errLog, func(c *gin.Context, err any) {
		// log := logger.ExtractLogger(c)
		// log.Error(errLog)
		errorutil.CastErrorJson(c, errormodel.RaiseTechnicalError("9999", "internal server error"))
		errLog.Reset()
		c.Next()
	})
}
