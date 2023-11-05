package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kumin/AndPadDating/erroz"
	"github.com/kumin/AndPadDating/handler"
	"github.com/kumin/AndPadDating/services"
)

func ValidateToken() gin.HandlerFunc {
	fmt.Println("hahahah")
	return func(c *gin.Context) {
		token := c.GetHeader("x-token")
		if !services.ValidateToken(token) {
			c.JSON(http.StatusBadRequest, handler.ErrorMessage(erroz.ErrBadToken))
			c.Abort()
			return
		}
		c.Next()
	}
}
