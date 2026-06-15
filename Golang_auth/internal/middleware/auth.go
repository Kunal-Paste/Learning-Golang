package middleware

import (
	"go-auth/internal/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	ctxUserIDkey = "auth.userId"
	ctxRoleKey   = "auth.role"
)

func AuthRequired(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := strings.TrimSpace(c.GetHeader("Authorization"))
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "missing Authorization token",
			})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid authorization event",
			})
			return
		}

		scheme := strings.TrimSpace(parts[0])
		tokenString := strings.TrimSpace(parts[1])

		if !strings.EqualFold(scheme, "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization scheme must be bearer",
			})
			return
		}

		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "missing token here",
			})
			return
		}

		claims, err := auth.ParseToken(jwtSecret, tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid or expired token",
			})
			return
		}

		c.Set(ctxUserIDkey, claims.Subject)
		c.Set(ctxRoleKey, claims.Role)

		c.Next()

	}
}

func GetUserID(c *gin.Context) (string, bool) {
	res, ok := c.Get(ctxUserIDkey)
	if !ok {
		return "", false
	}

	userID, ok := res.(string)
	return userID, ok
}

func GetRole(c *gin.Context) (string, bool) {
	res, ok := c.Get(ctxRoleKey)
	if !ok {
		return "", false
	}

	userRole, ok := res.(string)
	return userRole, ok
}
