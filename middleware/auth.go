package middleware

import (
    "net/http"
    "strings"
    "github.com/gin-gonic/gin"
    "ecommerce-api/utils"
)

func AuthMiddleware(role string) gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization header"})
            return
        }

        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        claims, err := utils.ValidateJWT(tokenString)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
            return
        }

        if role != "" && claims.Role != role {
            c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Unauthorized access"})
            return
        }

        c.Set("userID", claims.UserID)
        c.Set("role", claims.Role)
        c.Next()
    }
}
