package kernel

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v4/stdlib" // Import the pgx driver
	user_application "go-boilerplate/internal/app/module/user/application"
	user_domain "go-boilerplate/internal/app/module/user/domain"
	"go-boilerplate/internal/app/module/user/infrastructure"
	user_ui "go-boilerplate/internal/app/module/user/ui"
	"go-boilerplate/pkg/config"
	"go-boilerplate/pkg/http/middleware"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"net/http"
	"os"
)

const (
	UserSignInIndexRoute = "/sign-in"
	GoogleRedirectUrl    = "/auth/google"
	GoogleCallback       = "/auth/google/callback"

	GetUserList = "/users"
)

type UserModule struct {
	BaseModule

	UserSignInIndexHandler gin.HandlerFunc
	GoogleRedirectHandler  *user_ui.GoogleLoginRedirectHandler
	GoogleCallbackHandler  *user_ui.HandleGoogleCallbackHandler
	UserDashboardHandler   *user_ui.UserDashboardHandler

	GetUserList *user_ui.GetUserListHandler

	UserRepository user_domain.UserRepository
}

func (m *UserModule) Name() string {
	return "user_module"
}

// InitUserModule creates a new instance of NotificationModule.
func InitUserModule(k *Kernel, cnf *config.Config) *UserModule {
	_, err := user_infrastructure.NewPostgresUserRepository(cnf.DatabaseDSN)
	if err != nil {
		panic(err)
	}

	var googleOauthConfig = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_CLIENT_REDIRECT_URL"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}

	r, err := user_infrastructure.NewPostgresUserRepository(cnf.DatabaseDSN)
	if err != nil {
		panic("error connecting with the database")
	}

	um := &UserModule{
		UserRepository:         r,
		UserSignInIndexHandler: user_ui.HandleUserSocialSignInIndex,
		GoogleRedirectHandler:  user_ui.NewGoogleLoginRedirectHandler(googleOauthConfig),
		GoogleCallbackHandler:  user_ui.NewHandleGoogleCallbackHandler(googleOauthConfig, k.CommandBus),
		UserDashboardHandler:   user_ui.NewUserDashboardHandler(),
		GetUserList:            user_ui.NewGetUserListHandler(r),
	}

	um.AddCommand(&user_application.CreateUserCommand{}, user_application.NewCreateUserCommandHandler(r))

	return um
}

func (m *UserModule) RegisterRoutes(c *Kernel) {
	c.Router.Handle(
		http.MethodGet,
		UserSignInIndexRoute,
		m.UserSignInIndexHandler,
	)

	c.Router.Handle(
		http.MethodGet,
		GoogleRedirectUrl,
		m.GoogleRedirectHandler.HandleGoogleLoginRedirect,
	)

	c.Router.Handle(
		http.MethodGet,
		GoogleCallback,
		m.GoogleCallbackHandler.HandleGoogleCallback,
	)

	c.Router.WithMiddleware(middleware.AuthMiddleware).Handle(
		http.MethodGet,
		"/dashboard",
		m.UserDashboardHandler.HandleUserDashboard,
	)

	c.Router.WithMiddleware(middleware.AuthMiddleware).Handle(
		http.MethodGet,
		GetUserList,
		m.GetUserList.HandleGetUserList,
	)
}
