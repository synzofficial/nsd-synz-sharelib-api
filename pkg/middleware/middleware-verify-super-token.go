package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func MiddlewareVerifySuperToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			apiKey                   = c.GetHeader("X-API-KEY")
			keyHash                  = "$2a$14$s9fuxh38cO6kM8aEJSOfLOLBOgWxYcVPTXQ1Cj92HLDf/jTUj.4j2"
			isValidateAPIKeyComplete = false
		)

		if err := bcrypt.CompareHashAndPassword([]byte(keyHash), []byte(apiKey)); err == nil {
			isValidateAPIKeyComplete = true
		}

		httpCode, err := validateAuthorizationHeader(c)
		if err != nil {
			if isValidateAPIKeyComplete {
				c.Next()
				return
			}

			c.AbortWithStatusJSON(httpCode, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.Next()
	}
}

func validateAuthorizationHeader(c *gin.Context) (int, error) {
	token := c.GetHeader("Authorization")
	st := strings.Split(token, " ")
	if len(st) < 2 || st[0] != "Bearer" {
		return http.StatusUnauthorized, errors.New("unauthorized")
	}

	ht := &http.Client{
		Timeout: time.Second * 10,
	}

	url := os.Getenv("AUTH_HOST")

	req, err := http.NewRequestWithContext(c.Request.Context(), http.MethodGet, url, nil)
	if err != nil {
		return http.StatusInternalServerError, errors.New("unable to verify token")
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", fmt.Sprintf("%s%s", "sAccessToken=", st[1]))

	res, err := ht.Do(req)
	if err != nil {
		fmt.Printf("sharelib: unalbe to validate token with auth-api: %+v", err)
		return http.StatusInternalServerError, errors.New("unable to set validate token payload")
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return http.StatusInternalServerError, errors.New("validate token fail")
	}

	type payload struct {
		Message string `json:"message"`
		UserId  string `json:"userId"` // sub userid (now this is super token userId)
	}
	var p payload
	if err := json.Unmarshal(body, &p); err != nil {
		return http.StatusInternalServerError, errors.New("unable to extract authorize result")
	}

	if p.Message != "" {
		return http.StatusInternalServerError, fmt.Errorf("authorize service error: %+v", p.Message)
	}

	// set sub userId to context
	co := context.WithValue(c.Request.Context(), ctxKeySub, p.UserId)
	c.Request = c.Request.WithContext(co)

	return 0, nil
}
