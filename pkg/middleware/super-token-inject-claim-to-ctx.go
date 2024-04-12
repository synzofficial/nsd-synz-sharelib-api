package middleware

import (
	"context"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

// SuperTokenInjectClaimsToCtx: middleware for inject superTokenId to context with out authorize
func SuperTokenInjectClaimsToCtx() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		s := strings.Split(authorization, " ")
		if len(s) < 2 || s[0] != "Bearer" {
			c.Next()
			return
		}
		authorization = s[1]
		claimsMap, err := extracToken(authorization)
		if err != nil {
			fmt.Printf("SuperTokenInjectTokenIdOnly: unable to extract token -> %+v", err)
			c.Next()
			return
		}

		ctx := context.WithValue(c.Request.Context(), ctxKeySub, claimsMap["sub"])
		c.Request = c.Request.WithContext(ctx)
	}
}
