package user_ui

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"go-boilerplate/internal/app/module/user/ui/views"
	"log"
	"net/http"
)

type UserTestHandler struct {
}

func NewUserTestHandler() *UserTestHandler {
	return &UserTestHandler{}
}

func (uth *UserTestHandler) HandleUserTests(g *gin.Context) {
	_, exists := g.Get("userInfo")
	if !exists {
		log.Println("UserInfo not found in context")
		g.JSON(http.StatusInternalServerError, gin.H{"error": "User info not found"})
		return
	}

	g.Writer.Header().Set("Content-Type", "text/html")
	templ.Handler(views.DashboardLayout()).ServeHTTP(g.Writer, g.Request)
}
