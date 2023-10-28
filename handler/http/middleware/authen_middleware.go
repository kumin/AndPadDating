package middleware

import "github.com/gin-gonic/gin"

func ValidateLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO: validate token
	}
}
