package middleware

import (
	"go-boilerplate/pkg/router"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware ensures that the user is authenticated
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")

		if user == nil {
			// If user is not authenticated, respond with an error
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Cast the user to UserInfo and add it to the context
		userInfo, ok := user.(router.UserInfo)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid session data"})
			c.Abort()
			return
		}

		// Store UserInfo in the context
		c.Set("userInfo", userInfo)

		c.Next()
	}
}
