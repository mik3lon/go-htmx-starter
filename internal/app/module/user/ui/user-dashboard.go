package user_ui

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	user_domain "github.com/mik3lon/go-htmx-starter/internal/app/module/user/domain"
	"github.com/mik3lon/go-htmx-starter/internal/app/module/user/ui/views"
	"log"
	"net/http"
)

type UserDashboardHandler struct {
}

func NewUserDashboardHandler() *UserDashboardHandler {
	return &UserDashboardHandler{}
}

func (uth *UserDashboardHandler) HandleUserDashboard(g *gin.Context) {
	u, exists := g.Get("user")
	if !exists {
		log.Println("UserInfo not found in context")
		g.JSON(http.StatusInternalServerError, gin.H{"error": "User info not found"})
		return
	}

	user, ok := u.(user_domain.User)
	if !ok {
		panic("user not found")
	}

	g.Writer.Header().Set("Content-Type", "text/html")
	templ.Handler(views.DashboardLayout(user)).ServeHTTP(g.Writer, g.Request)
}
