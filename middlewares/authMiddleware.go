package middlewares

import (
	"golang-resturant-management/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")

		if token == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "no authentication token provided"})
			c.Abort()
			return
		}

		claims, msg := helpers.ValidateToken(token)
		if msg != "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": msg})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("first_name", claims.First_name)
		c.Set("last_name", claims.Last_name)
		c.Set("uid", claims.Uid)

		c.Next()
	}
}
