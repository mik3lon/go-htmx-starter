package user_ui

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/mik3lon/go-htmx-starter/internal/app/module/user/ui/views"
)

func HandleUserSocialSignInIndex(g *gin.Context) {
	g.Writer.Header().Set("Content-Type", "text/html")
	templ.Handler(views.Index()).ServeHTTP(g.Writer, g.Request)
}
