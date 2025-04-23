package utils

import (
        "github.com/gin-gonic/gin"
)

// ErrorResponse sends a standardized error response
func ErrorResponse(ctx *gin.Context, statusCode int, message string, err error) {
        details := ""
        if err != nil {
                details = err.Error()
        }

        ctx.JSON(statusCode, gin.H{
                "error": gin.H{
                        "message": message,
                        "details": details,
                },
        })
}
