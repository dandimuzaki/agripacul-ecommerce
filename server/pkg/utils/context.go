package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func GetUserIDFromContext(c *gin.Context) (uint, error) {
	id, exists := c.Get("user_id")
	if !exists {
		return 0, errors.New("user ID not found in context")
	}

	switch v := id.(type) {
	case uint:
		return v, nil
	case float64:
		return uint(v), nil
	default:
		return 0, errors.New("invalid user ID type")
	}
}

func GetRoleFromContext(c *gin.Context) (string, error) {
	role, exists := c.Get("role")
	if !exists {
		return "", errors.New("role not found in context")
	}

	if roleStr, ok := role.(string); ok {
		return roleStr, nil
	}

	return "", errors.New("invalid role type")
}
