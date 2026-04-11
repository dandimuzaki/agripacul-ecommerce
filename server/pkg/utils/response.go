package utils

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Success  bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Errors  any    `json:"errors,omitempty"`
	Fields  any    `json:"fields,omitempty"`
}

func ResponseSuccess(c *gin.Context, code int, message string, data any) {
	response := Response{
		Success:  true,
		Message: message,
		Data:    data,
	}
	c.JSON(code, response)
}

func ResponseFailed(c *gin.Context, code int, message string, errors any) {
	response := Response{
		Success:  false,
		Message: message,
		Errors:  errors,
	}
	c.JSON(code, response)
}

func ResponsePagination(c *gin.Context, code int, message string, data any, meta interface{}) {
	response := map[string]interface{}{
		"success":     true,
		"message":    message,
		"data":       data,
		"pagination": meta,
	}
	c.JSON(code, response)
}
