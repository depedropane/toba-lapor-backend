package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"toba-lapor-backend/pkg/utils"
)

func AuthMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.BuildErrorResponse(c, http.StatusUnauthorized, "Unauthorized", "No token provided")
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.BuildErrorResponse(c, http.StatusUnauthorized, "Unauthorized", "Invalid token format")
			return
		}

		tokenString := parts[1]
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			utils.BuildErrorResponse(c, http.StatusUnauthorized, "Unauthorized", "Invalid or expired token")
			return
		}

		// Check roles
		if len(roles) > 0 {
			validRole := false
			for _, role := range roles {
				if claims.RoleName == role {
					validRole = true
					break
				}
			}
			if !validRole {
				utils.BuildErrorResponse(c, http.StatusForbidden, "Forbidden", "You don't have permission to access this resource")
				return
			}
		}

		// Set context variable
		c.Set("user_id", claims.UserID)
		c.Set("role", claims.RoleName)
		if claims.AgencyID != nil {
			c.Set("agency_id", *claims.AgencyID)
		}

		c.Next()
	}
}
