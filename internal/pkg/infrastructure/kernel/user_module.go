package kernel

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v4/stdlib" // Import the pgx driver
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
)

type UserModule struct {
	UserSignInIndexHandler gin.HandlerFunc
	GoogleRedirectHandler  *user_ui.GoogleLoginRedirectHandler
	GoogleCallbackHandler  *user_ui.HandleGoogleCallbackHandler
	HandleUserTests        *user_ui.UserTestHandler
}

// InitUserModule creates a new instance of UserModule.
func InitUserModule(c *Kernel, cnf *config.Config) *UserModule {
	_, err := user_infrastructure.NewPostgresUserRepository(cnf.DatabaseDSN)
	if err != nil {
		panic(err)
	}

	var googleOauthConfig = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_CLIENT_REDIRECT_URL"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	um := &UserModule{
		UserSignInIndexHandler: user_ui.HandleUserSocialSignInIndex,
		GoogleRedirectHandler:  user_ui.NewGoogleLoginRedirectHandler(googleOauthConfig),
		GoogleCallbackHandler:  user_ui.NewHandleGoogleCallbackHandler(googleOauthConfig),
		HandleUserTests:        user_ui.NewUserTestHandler(),
	}

	return um
}

func (m *UserModule) RegisterRoutes(c *Kernel) {
	c.Router.Handle(
		http.MethodGet,
		UserSignInIndexRoute,
		m.UserSignInIndexHandler,
	)

	c.Router.WithMiddleware().Handle(
		http.MethodGet,
		GoogleRedirectUrl,
		m.GoogleRedirectHandler.HandleGoogleLoginRedirect,
	)

	c.Router.WithMiddleware().Handle(
		http.MethodGet,
		GoogleCallback,
		m.GoogleCallbackHandler.HandleGoogleCallback,
	)

	c.Router.WithMiddleware(middleware.AuthMiddleware).Handle(
		http.MethodGet,
		"/users/test",
		m.HandleUserTests.HandleUserTests,
	)
}
