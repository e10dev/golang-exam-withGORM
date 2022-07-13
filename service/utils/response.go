package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	OK                    = http.StatusOK
	CREATED               = http.StatusCreated
	NO_CONTENT            = http.StatusNoContent
	BAD_REQUEST           = http.StatusBadRequest
	UNAUTHORIZED          = http.StatusUnauthorized
	FORBIDDEN             = http.StatusForbidden
	NOT_FOUND             = http.StatusNotFound
	INTERNAL_SERVER_ERROR = http.StatusInternalServerError
	SERVICE_UNAVAILABLE   = http.StatusServiceUnavailable
)

func Response(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": msg,
	})
}

func ResponseJSON(c *gin.Context, jsonData interface{}) {
	c.JSON(OK, jsonData)
}

func ResponseCSVFile(c *gin.Context, data []byte) {
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename=users.csv")
	c.Data(OK, "text/csv", data)
}
