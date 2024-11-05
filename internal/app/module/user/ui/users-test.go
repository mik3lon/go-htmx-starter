package user_ui

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserTestHandler struct {
}

func NewUserTestHandler() *UserTestHandler {
	return &UserTestHandler{}
}

func (uth *UserTestHandler) HandleUserTests(g *gin.Context) {
	userInfo, exists := g.Get("userInfo")
	if !exists {
		log.Println("UserInfo not found in context")
		g.JSON(http.StatusInternalServerError, gin.H{"error": "User info not found"})
		return
	}

	response := map[string]interface{}{
		"message": "HERE!",
		"user":    userInfo,
	}
	if err := json.NewEncoder(g.Writer).Encode(response); err != nil {
		log.Printf("Failed to write response: %v\n", err)
	}

}
