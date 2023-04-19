package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/herizal95/project-starter/utils"
)

func DeserializerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// ambil token dari header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		// pisahkan token dengan prefix "Bearer "
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		// validasi token
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		// simpan data user di context untuk digunakan di handler selanjutnya
		c.Set("user", claims.User)

		c.Next()
	}
}
