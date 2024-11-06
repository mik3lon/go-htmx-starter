package user_ui

import (
	"github.com/gin-gonic/gin"
	"go-boilerplate/internal/app/module/user/ui/views"
)

type UserSettingsHandler struct {
}

func NewUserSettingsHandler() *UserSettingsHandler {
	return &UserSettingsHandler{}
}

func (uth *UserSettingsHandler) HandeUserSettings(g *gin.Context) {
	views.SettingsPage().Render(g, g.Writer)
}
