package notification_ui

import (
	"github.com/gin-gonic/gin"
	"github.com/mik3lon/go-htmx-starter/internal/app/module/notification/ui/views"
)

type GetNotificationListHandler struct {
}

func NewGetNotificationListHandler() *GetNotificationListHandler {
	return &GetNotificationListHandler{}
}

func (uth *GetNotificationListHandler) HandleGetUserList(g *gin.Context) {
	views.NotificationList().Render(g, g.Writer)
}
