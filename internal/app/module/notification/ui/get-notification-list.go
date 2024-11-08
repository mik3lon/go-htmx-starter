package notification_ui

import (
	"github.com/gin-gonic/gin"
	"go-boilerplate/internal/app/module/notification/ui/views"
)

type GetNotificationListHandler struct {
}

func NewGetNotificationListHandler() *GetNotificationListHandler {
	return &GetNotificationListHandler{}
}

func (uth *GetNotificationListHandler) HandleGetUserList(g *gin.Context) {
	views.NotificationList().Render(g, g.Writer)
}
