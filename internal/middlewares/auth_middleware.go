package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"hospital-portal/internal/auth"
	"hospital-portal/internal/utils"
)

// AuthMiddleware ensures that requests are authenticated
func AuthMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get the Authorization header
		authHeader := ctx.GetHeader("Authorization")
		
		// Check if Authorization header exists and has the correct format
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			logger.Warn("Missing or invalid Authorization header")
			utils.ErrorResponse(ctx, http.StatusUnauthorized, "Authorization header is required", nil)
			ctx.Abort()
			return
		}

		// Extract the token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		
		// Validate the token
		claims, err := auth.ValidateToken(tokenString)
		if err != nil {
			logger.Warn("Invalid token", zap.Error(err))
			utils.ErrorResponse(ctx, http.StatusUnauthorized, "Invalid or expired token", err)
			ctx.Abort()
			return
		}
		
		// Store user information in the context
		ctx.Set("user_id", claims.UserID)
		ctx.Set("user_email", claims.Email)
		ctx.Set("user_role", claims.Role)
		
		ctx.Next()
	}
}

// RoleMiddleware checks if the user has the required role
func RoleMiddleware(roles ...auth.Role) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get user role from context
		userRole, exists := ctx.Get("user_role")
		if !exists {
			utils.ErrorResponse(ctx, http.StatusUnauthorized, "Unauthorized", nil)
			ctx.Abort()
			return
		}

		// Check if user has one of the required roles
		hasRequiredRole := false
		for _, role := range roles {
			if userRole == role {
				hasRequiredRole = true
				break
			}
		}

		if !hasRequiredRole {
			utils.ErrorResponse(ctx, http.StatusForbidden, "Access denied", nil)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
