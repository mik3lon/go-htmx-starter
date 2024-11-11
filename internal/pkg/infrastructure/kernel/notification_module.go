package kernel

import (
	_ "github.com/jackc/pgx/v4/stdlib" // Import the pgx driver
	notification_ui "go-boilerplate/internal/app/module/notification/ui"
	"go-boilerplate/pkg/bus"
	"go-boilerplate/pkg/bus/command"
	"go-boilerplate/pkg/config"
	"net/http"
)

const (
	GetNotificationList = "/notifications"
)

type NotificationModule struct {
	GetNotificationList *notification_ui.GetNotificationListHandler
}

func (m *NotificationModule) Name() string {
	return "notification_module"
}

func InitNotificationModule(c *Kernel, cnf *config.Config) *NotificationModule {

	um := &NotificationModule{
		GetNotificationList: notification_ui.NewGetNotificationListHandler(),
	}

	return um
}

func (m *NotificationModule) RegisterRoutes(c *Kernel) {
	c.Router.Handle(
		http.MethodGet,
		GetNotificationList,
		m.GetNotificationList.HandleGetUserList,
	)
}

func (m *NotificationModule) Commands() map[bus.Dto]command.CommandHandler {
	return map[bus.Dto]command.CommandHandler{}
}
