package user_ui

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"go-boilerplate/internal/app/module/user/ui/views"
)

func HandleUserSocialSignInIndex(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/html")
	templ.Handler(views.Index()).ServeHTTP(c.Writer, c.Request)
}
