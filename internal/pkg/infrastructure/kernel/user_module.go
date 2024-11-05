package kernel

import (
	"go-boilerplate/internal/app/module/user/infrastructure"
	user_ui "go-boilerplate/internal/app/module/user/ui"
	"go-boilerplate/pkg/config"
	"go-boilerplate/pkg/http/middleware"
	"net/http"

	_ "github.com/jackc/pgx/v4/stdlib" // Import the pgx driver
)

type UserModule struct {
}

// InitUserModule creates a new instance of UserModule.
func InitUserModule(c *Kernel, cnf *config.Config) *UserModule {
	_, err := user_infrastructure.NewPostgresUserRepository(cnf.DatabaseDSN)
	if err != nil {
		panic(err)
	}

	um := &UserModule{}

	return um
}

// RegisterRoutes registers the user module routes.
func (m *UserModule) RegisterRoutes(c *Kernel) {
	c.Router.WithMiddleware(middleware.LoggingMiddleware).Handle(
		http.MethodGet,
		"/sign-in",
		user_ui.HandleUserSocialSignInIndex,
	)

	c.Router.WithMiddleware(middleware.LoggingMiddleware).Handle(
		http.MethodGet,
		"/auth/google",
		user_ui.HandleGoogleLoginRedirect,
	)

	c.Router.WithMiddleware(middleware.LoggingMiddleware).Handle(
		http.MethodGet,
		"/auth/google/callback",
		user_ui.HandleGoogleCallback,
	)
}
