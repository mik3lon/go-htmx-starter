package user_application

import (
	"context"
	"errors"
	user_domain "github.com/mik3lon/go-htmx-starter/internal/app/module/user/domain"
	"github.com/mik3lon/go-htmx-starter/pkg/bus"
)

type CreateUserCommand struct {
	ID                string
	Name              string
	Surname           string
	Username          string
	PlainPassword     string
	Email             string
	Role              string
	ProfilePictureUrl string
	IsFormSocialAuth  bool
}

func (c CreateUserCommand) Id() string {
	return "create-user-command"
}

type CreateUserCommandHandler struct {
	r user_domain.UserRepository
}

func NewCreateUserCommandHandler(r user_domain.UserRepository) *CreateUserCommandHandler {
	return &CreateUserCommandHandler{r: r}
}

func (cuch CreateUserCommandHandler) Handle(ctx context.Context, c bus.Dto) error {
	cuc, ok := c.(*CreateUserCommand)
	if !ok {
		return errors.New("invalid command")
	}

	password, err := user_domain.GenerateHashedPassword(cuc.IsFormSocialAuth, cuc.PlainPassword)
	if err != nil {
		return errors.New("failed to generate hashed password")
	}

	user := user_domain.CreateUser(
		cuc.ID,
		cuc.Username,
		cuc.Email,
		password,
		cuc.Name,
		cuc.Surname,
		cuc.Role,
		cuc.ProfilePictureUrl,
	)

	return cuch.r.Save(ctx, user)
}
