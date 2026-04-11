package middleware

import (
	"debian-ecommerce/internal/data/repository"
	"debian-ecommerce/internal/usecase"
	"debian-ecommerce/pkg/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MiddlewareCustom struct {
	Usecase      *usecase.Usecase
	Log          *zap.Logger
	TokenService utils.TokenService
	TokenRepo    repository.TokenRepository
}

func NewMiddlewareCustom(usecase *usecase.Usecase, log *zap.Logger, tokenService utils.TokenService, tokenRepo repository.TokenRepository) MiddlewareCustom {
	return MiddlewareCustom{
		Usecase:      usecase,
		Log:          log,
		TokenService: tokenService,
		TokenRepo:    tokenRepo,
	}
}

func (m *MiddlewareCustom) BearerTokenAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			utils.ResponseFailed(ctx, http.StatusUnauthorized, "unauthorized", nil)
			ctx.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.ResponseFailed(ctx, http.StatusUnauthorized, "invalid token format", nil)
			ctx.Abort()
			return
		}

		tokenString := parts[1]
		claims, err := m.TokenService.ValidateToken(tokenString)
		if err != nil {
			utils.ResponseFailed(ctx, http.StatusUnauthorized, "invalid token", err.Error())
			ctx.Abort()
			return
		}

		ctx.Set("user_id", claims.UserID)
		ctx.Set("role", claims.Role)

		// Validate against Redis
		if err := m.TokenRepo.ValidateToken(ctx.Request.Context(), claims.UserID, tokenString); err != nil {
			utils.ResponseFailed(ctx, http.StatusUnauthorized, "session expired or invalid", err.Error())
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func (m *MiddlewareCustom) RBAC(allowedRoles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role, exists := ctx.Get("role")
		if !exists {
			utils.ResponseFailed(ctx, http.StatusUnauthorized, "unauthorized", nil)
			ctx.Abort()
			return
		}

		roleStr, ok := role.(string)
		if !ok {
			utils.ResponseFailed(ctx, http.StatusInternalServerError, "invalid role type", nil)
			ctx.Abort()
			return
		}

		isAllowed := false
		for _, allowedRole := range allowedRoles {
			if roleStr == allowedRole {
				isAllowed = true
				break
			}
		}

		if !isAllowed {
			utils.ResponseFailed(ctx, http.StatusForbidden, "forbidden", nil)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
