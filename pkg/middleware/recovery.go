package middleware

import (
	"bytes"

	"github.com/gin-gonic/gin"
	"github.com/synzofficial/nsd-synz-sharelib-api/pkg/model/errormodel"
	"github.com/synzofficial/nsd-synz-sharelib-api/pkg/util/errorutil"
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
