package ginutil

import (
	"bytes"

	"github.com/gin-gonic/gin"
	errorutil "github.com/synzofficial/nsd-synz-sharelib-api/pkg/util/error-util"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func CastErrorJsonWithLogging(c *gin.Context, err error) {
	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw

	errorutil.CastErrorJson(c, err)
}
