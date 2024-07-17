package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"service-manager/http/authorization"
)

func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		token := authorization.NewToken()
		_, err := token.Analyze(auth)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Next()
	}
}
