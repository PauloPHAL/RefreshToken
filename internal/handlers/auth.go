package handlers

import (
	"net/http"

	"github.com/PauloPHAL/microservices/internal/views"
	"github.com/PauloPHAL/microservices/pkg/dto"
	"github.com/PauloPHAL/microservices/pkg/interfaces"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service interfaces.AuthService
}

func NewAuthHandler(service interfaces.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var login dto.LoginDTO

	if err := c.BindJSON(&login); err != nil {
		views.SendError(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := login.Validate(); err != nil {
		views.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.service.Login(c.Request.Context(), &login)
	if err != nil {
		views.SendError(c, http.StatusUnauthorized, err.Error())
		return
	}

	views.SendSuccess(c, http.StatusOK, response)
}

func (h *AuthHandler) Refresh(c *gin.Context) {
	var request dto.RefreshTokenDTO

	if err := c.BindJSON(&request); err != nil {
		views.SendError(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := request.Validate(); err != nil {
		views.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.service.Refresh(c.Request.Context(), &request)
	if err != nil {
		views.SendError(c, http.StatusUnauthorized, err.Error())
		return
	}

	views.SendSuccess(c, http.StatusOK, response)
}

func (h *AuthHandler) AuthJWT(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		views.SendError(c, http.StatusUnauthorized, "Authorization header required")
		c.Abort()
		return
	}

	userID, err := h.service.ValidateToken(c.Request.Context(), tokenString)
	if err != nil {
		views.SendError(c, http.StatusUnauthorized, err.Error())
		c.Abort()
		return
	}

	c.Set("userID", userID)
	c.Next()
}

func (h *AuthHandler) ValidateToken(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		views.SendError(c, http.StatusUnauthorized, "User ID not found in context")
		return
	}

	views.SendSuccess(c, http.StatusOK, gin.H{
		"message": "Token is valid",
		"userID":  userID,
	})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		views.SendError(c, http.StatusUnauthorized, "User ID not found in context")
		return
	}

	if err := h.service.Logout(c.Request.Context(), userID.(string)); err != nil {
		views.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	views.SendSuccess(c, http.StatusOK, gin.H{"message": "Logged out successfully"})
}
