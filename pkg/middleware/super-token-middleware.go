package middleware

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/session/claims"
)

var (
	ctxKeySub = struct{}{}
)

func GetSuperTokenUserID(ctx context.Context) string {
	sub, ok := ctx.Value(ctxKeySub).(string)
	if !ok {
		return ""
	}
	return sub
}

func SuperTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authorization := c.GetHeader("Authorization")
		s := strings.Split(authorization, " ")
		if len(s) != 2 {
			c.AbortWithStatus(401)
			return
		}
		authorization = s[1]
		claimsMap, err := extracToken(authorization)
		if err != nil {
			c.AbortWithStatus(401)
			return
		}
		co := context.WithValue(c.Request.Context(), ctxKeySub, claimsMap["sub"])
		c.Request = c.Request.WithContext(co)

		// ------
		// TODO: handle error
		claim, err := session.GetClaimValue(authorization, &claims.TypeSessionClaim{})
		_, _ = claim, err
		// ------

		c.Next()
	}
}

func SuperTokenWithNotVerifyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authorization := c.GetHeader("Authorization")
		s := strings.Split(authorization, " ")
		if len(s) != 2 {
			c.Next()
			return
		}
		authorization = s[1]
		claimsMap, err := extracToken(authorization)
		if err != nil {
			c.AbortWithStatus(401)
			return
		}
		co := context.WithValue(c.Request.Context(), ctxKeySub, claimsMap["sub"])
		c.Request = c.Request.WithContext(co)

		// ------
		// TODO: handle error
		claim, err := session.GetClaimValue(authorization, &claims.TypeSessionClaim{})
		_, _ = claim, err
		// ------

		c.Next()
	}
}

func extracToken(tokenString string) (jwt.MapClaims, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		fmt.Printf("Error %s", err)
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("unalbe to extract claims token")
	}
	return claims, nil
}
