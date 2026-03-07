package handlers

import (
	"net/http"

	"github.com/PauloPHAL/microservices/internal/views"
	"github.com/PauloPHAL/microservices/pkg/dto"
	"github.com/PauloPHAL/microservices/pkg/interfaces"
	"github.com/PauloPHAL/microservices/pkg/perrors"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service interfaces.UserService
}

func NewUserHandler(service interfaces.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (u *UserHandler) CreateUser(c *gin.Context) {
	var user dto.UserDTO

	if err := c.ShouldBindJSON(&user); err != nil {
		views.SendError(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := user.Validate(); err != nil {
		views.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	response, err := u.service.CreateUser(c.Request.Context(), &user)
	if err != nil {
		code := http.StatusInternalServerError
		if err.Error() == perrors.ErrEmailAlreadyExists.Error() {
			code = http.StatusConflict
		}
		views.SendError(c, code, err.Error())
		return
	}

	views.SendSuccess(c, http.StatusCreated, response)
}

func (u *UserHandler) GetUser(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		views.SendError(c, http.StatusBadRequest, "User ID is required")
		return
	}

	user, err := u.service.GetUser(c.Request.Context(), id)
	if err != nil {
		code := http.StatusInternalServerError
		if err.Error() == perrors.ErrUserNotFound.Error() {
			code = http.StatusNotFound
		}
		views.SendError(c, code, err.Error())
		return
	}

	views.SendSuccess(c, http.StatusOK, user)
}
