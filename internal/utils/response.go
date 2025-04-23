package utils

import (
	"github.com/gin-gonic/gin"
)

// Response is a generic response structure
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// SuccessResponse sends a standardized success response
func SuccessResponse(ctx *gin.Context, statusCode int, message string, data interface{}) {
	ctx.JSON(statusCode, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// PaginatedResponse represents a paginated response
type PaginatedResponse struct {
	Items      interface{} `json:"items"`
	TotalItems int64       `json:"totalItems"`
	Page       int         `json:"page"`
	PageSize   int         `json:"pageSize"`
	TotalPages int         `json:"totalPages"`
}

// PaginateResponse sends a standardized paginated response
func PaginateResponse(ctx *gin.Context, statusCode int, items interface{}, totalItems int64, page, pageSize int) {
	totalPages := int(totalItems) / pageSize
	if int(totalItems)%pageSize > 0 {
		totalPages++
	}

	ctx.JSON(statusCode, gin.H{
		"success": true,
		"data": PaginatedResponse{
			Items:      items,
			TotalItems: totalItems,
			Page:       page,
			PageSize:   pageSize,
			TotalPages: totalPages,
		},
	})
}
