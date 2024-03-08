package ginutil

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/synzofficial/nsd-synz-sharelib-api/pkg/middleware"
)

// NewEngie: initial gin engine with default middleware{CORS, ExtractJWT, Recover}
func NewEngie() *gin.Engine {
	mode := os.Getenv("GIN_MODE")
	switch mode {
	case "release", "debug", "test":
	default:
		mode = "release"
	}
	gin.SetMode(mode)

	engine := gin.New()
	engine.Use(middleware.CORS())
	engine.Use(middleware.Recovery())
	return engine
}
