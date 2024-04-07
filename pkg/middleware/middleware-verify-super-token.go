package middleware

import (
	"encoding/json"
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
		apiKey := c.GetHeader("X-API-KEY")

		keyHash := "$2a$14$s9fuxh38cO6kM8aEJSOfLOLBOgWxYcVPTXQ1Cj92HLDf/jTUj.4j2"

		if err := bcrypt.CompareHashAndPassword([]byte(keyHash), []byte(apiKey)); err == nil {
			c.Next()
			return
		}

		token := c.GetHeader("Authorization")
		st := strings.Split(token, " ")
		if len(st) < 2 || st[0] != "Bearer" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ht := &http.Client{
			Timeout: time.Second * 10,
		}

		url := os.Getenv("AUTH_HOST")

		req, err := http.NewRequestWithContext(c.Request.Context(), http.MethodGet, url, nil)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "unable to verify token",
			})
			return
		}
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Cookie", fmt.Sprintf("%s%s", "sAccessToken=", st[1]))

		res, err := ht.Do(req)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "validate token fail",
			})
			return
		}
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "validate token fail",
			})
			return
		}

		type payload struct {
			Message string `json:"message"`
		}
		var p payload
		if err := json.Unmarshal(body, &p); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "validate token fail",
			})
			return
		}

		if p.Message != "" {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": p.Message,
			})
			return
		}

		c.Next()
	}
}
