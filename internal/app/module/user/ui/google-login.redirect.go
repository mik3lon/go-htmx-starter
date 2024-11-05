package user_ui

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"net/http"
)

type GoogleLoginRedirectHandler struct {
	googleOauthConfig *oauth2.Config
}

func NewGoogleLoginRedirectHandler(googleOauthConfig *oauth2.Config) *GoogleLoginRedirectHandler {
	return &GoogleLoginRedirectHandler{googleOauthConfig: googleOauthConfig}
}

func (h *GoogleLoginRedirectHandler) HandleGoogleLoginRedirect(g *gin.Context) {
	url := h.googleOauthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	g.Redirect(http.StatusFound, url)
}
