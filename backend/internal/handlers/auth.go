package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"social-media-scheduler/internal/dtos"
	"social-media-scheduler/internal/services"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var request dtos.LoginRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session, err := h.authService.Login(c.Request.Context(), request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("session_id", session.ID.String(), 3600*24, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func (h *AuthHandler) Register(c *gin.Context) {
	var request dtos.RegisterRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.authService.Register(c.Request.Context(), request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Signup successful"})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	userIdStr, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	err := h.authService.Logout(c.Request.Context(), userIdStr.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("session_id", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
