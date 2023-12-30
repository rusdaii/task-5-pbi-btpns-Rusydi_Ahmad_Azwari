package middlewares

import (
	"go-project/app"
	"go-project/helpers"
	"go-project/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.GetHeader("Authorization")

		splitToken := strings.Split(bearerToken, "Bearer ")

		if len(splitToken) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
			c.Abort()
			return
		}

		token := splitToken[1]

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			c.Abort()
			return
		}

		claims, err := helpers.ValidateToken(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}

		user := &models.User{
			Base: app.Base{
				ID: claims.ID,
			},
		}

		c.Set("user", user)

		c.Next()
	}
}
