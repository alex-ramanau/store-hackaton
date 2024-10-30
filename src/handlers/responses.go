package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func JSONOkResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   data,
	})
}

func JSONErrorResponse(c *gin.Context, code int, errorCode string, errorMessage string, context interface{}) {
	c.JSON(code, gin.H{
		"status":        "error",
		"error_code":    errorCode,
		"error_message": errorMessage,
		"context":       context,
	})
}

