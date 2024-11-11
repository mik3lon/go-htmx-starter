package user_ui

import (
	"github.com/gin-gonic/gin"
	user_application "github.com/mik3lon/go-htmx-starter/internal/app/module/user/application"
	user_domain "github.com/mik3lon/go-htmx-starter/internal/app/module/user/domain"
	"github.com/mik3lon/go-htmx-starter/internal/app/module/user/ui/views"
	"github.com/mik3lon/go-htmx-starter/pkg/bus/query"
	"strconv"
)

type GetUserListHandler struct {
	qb query.Bus
}

func NewGetUserListHandler(qb query.Bus) *GetUserListHandler {
	return &GetUserListHandler{qb: qb}
}

func (uth *GetUserListHandler) HandleGetUserList(g *gin.Context) {
	page, err := strconv.Atoi(g.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	size, err := strconv.Atoi(g.DefaultQuery("size", "20"))
	if err != nil || size < 1 {
		size = 20
	}

	userList, err := uth.qb.Ask(g, &user_application.ListUsersQuery{
		Page:   page,
		Size:   size,
		Filter: nil,
	})
	if err != nil {
		panic(err)
	}

	views.UserList(userList.(user_domain.UserList), page, size, 100).Render(g, g.Writer)
}
