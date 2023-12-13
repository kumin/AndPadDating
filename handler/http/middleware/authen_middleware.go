package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kumin/BityDating/erroz"
	"github.com/kumin/BityDating/handler"
	"github.com/kumin/BityDating/services"
)

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := strings.Split(c.GetHeader("authorization"), " ")
		if len(authToken) < 2 {
			c.JSON(http.StatusBadRequest, handler.ErrorMessage(erroz.ErrBadToken))
			c.Abort()
			return
		}
		if !services.ValidateToken(authToken[1]) {
			c.JSON(http.StatusBadRequest, handler.ErrorMessage(erroz.ErrBadToken))
			c.Abort()
			return
		}
		c.Next()
	}
}
