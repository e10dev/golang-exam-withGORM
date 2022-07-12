package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": msg,
	})
}

func ResponseJSON(c *gin.Context, jsonData interface{}) {
	c.JSON(http.StatusOK, jsonData)
}
