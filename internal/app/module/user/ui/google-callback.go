package user_ui

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	user_application "go-boilerplate/internal/app/module/user/application"
	user_domain "go-boilerplate/internal/app/module/user/domain"
	"go-boilerplate/pkg/bus/command"
	"go-boilerplate/pkg/router"
	"golang.org/x/oauth2"
	"log"
	"net/http"
)

type HandleGoogleCallbackHandler struct {
	googleOauthConfig *oauth2.Config
	cb                *command.CommandBus
}

func NewHandleGoogleCallbackHandler(googleOauthConfig *oauth2.Config, cb *command.CommandBus) *HandleGoogleCallbackHandler {
	return &HandleGoogleCallbackHandler{googleOauthConfig: googleOauthConfig, cb: cb}
}

func (h *HandleGoogleCallbackHandler) HandleGoogleCallback(g *gin.Context) {
	code := g.Query("code")
	if code == "" {
		http.Error(g.Writer, `{"error": "Code not found"}`, http.StatusBadRequest)
		return
	}

	// Exchange the authorization code for an access token
	token, err := h.googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		log.Printf("Failed to exchange token: %v\n", err)
		http.Error(g.Writer, `{"error": "Failed to exchange token"}`, http.StatusInternalServerError)
		return
	}

	// Use the access token to get user info
	client := h.googleOauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		log.Printf("Failed to get user info: %v\n", err)
		http.Error(g.Writer, `{"error": "Failed to get user info"}`, http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Display user info as JSON (for demonstration purposes)
	var userInfo router.GoogleUserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		log.Printf("Failed to parse user info: %v\n", err)
		http.Error(g.Writer, `{"error": "Failed to parse user info"}`, http.StatusInternalServerError)
		return
	}

	cuc := &user_application.CreateUserCommand{
		ID:               uuid.NewString(),
		Name:             userInfo.GivenName,
		Surname:          userInfo.FamilyName,
		Username:         userInfo.Name,
		PlainPassword:    "",
		Email:            userInfo.Email,
		Role:             "ROL_USER",
		IsFormSocialAuth: true,
	}

	err = h.cb.Dispatch(g, cuc)
	var userAlreadyExists *user_domain.UserAlreadyExists
	switch {
	case errors.As(err, &userAlreadyExists):
		session := sessions.Default(g)
		session.Set("user", userInfo)
		if err := session.Save(); err != nil {
			log.Printf("Failed to save session: %v\n", err)
			http.Error(g.Writer, `{"error": "Failed to save session"}`, http.StatusInternalServerError)
			return
		}

		http.Redirect(g.Writer, g.Request, "/dashboard", http.StatusFound)
	default:
		http.Error(g.Writer, `{"error": "Failed to get user info"}`, http.StatusInternalServerError)
	}
}
