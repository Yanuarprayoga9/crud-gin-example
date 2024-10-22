package utils

import "github.com/gin-gonic/gin"

// SuccessResponse formats successful responses with a dynamic data key
func SuccessResponse(c *gin.Context, message string, dataKey string, data interface{}, httpStatusCode int) {
	c.JSON(httpStatusCode, gin.H{
		"message": message,
		"status":  "success",
		dataKey:   data,
	})
}

// ErrorResponse formats error responses (simplified version)
func ErrorResponse(c *gin.Context, message string, httpStatusCode int) {
	c.JSON(httpStatusCode, gin.H{
		"message": message,
		"status":  "failed",
	})
}
