package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kumin/BityDating/entities"
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
		claims, isValid := services.ValidateToken(authToken[1])
		if !isValid {
			c.JSON(http.StatusBadRequest, handler.ErrorMessage(erroz.ErrBadToken))
			c.Abort()
			return
		}
		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, entities.CtxUserIdKey, claims.UserId)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
