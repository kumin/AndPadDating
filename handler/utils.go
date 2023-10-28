package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ErrorMessage(err error) gin.H {
	return gin.H{
		"message": err.Error(),
	}
}

func GetParam(key string, params *gin.Params) (string, error) {
	val, ok := params.Get(key)
	if ok {
		return val, nil
	}
	return "", fmt.Errorf("%s is missing", key)
}
