package user_ui

import (
	"github.com/gin-gonic/gin"
	user_domain "go-boilerplate/internal/app/module/user/domain"
	"go-boilerplate/internal/app/module/user/ui/views"
)

type GetUserListHandler struct {
	r user_domain.UserRepository
}

func NewGetUserListHandler(r user_domain.UserRepository) *GetUserListHandler {
	return &GetUserListHandler{r: r}
}

func (uth *GetUserListHandler) HandleGetUserList(g *gin.Context) {
	userList, err := uth.r.FindAll(g)
	if err != nil {
		panic(err)
	}

	views.UserList(userList).Render(g, g.Writer)
}
