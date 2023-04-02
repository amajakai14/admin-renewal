package http

import (
	"context"
	"encoding/json"
	"net/http"

	appUser "github.com/amajakai14/admin-renewal/internal/user"
	"github.com/amajakai14/admin-renewal/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type UserService interface {
	PostUser(ctx context.Context, user appUser.User) (appUser.User, error)
	GetUser(ctx context.Context, id string) (appUser.User, error)
	UpdateUser(ctx context.Context, user appUser.User, id string) (appUser.User, error)
	DeleteUser(ctx context.Context, id string) error
}

type PostUserRequest struct {
	Name          string `json:"name"`
	Email         string `json:"email" validate:"required,email"`
	Password      string `json:"password" validate:"required,min=8"`
	Role          string `json:"role" validate:"required, oneof=admin staff"`
	CorporationId string `json:"corporation_id"`
}

func (h *Handler) PostUser(c *gin.Context) {
	r := c.Request
	w := c.Writer
	var userRequest PostUserRequest

	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	validate := validator.New()
	err := validate.Struct(userRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	user, err := toUser(userRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	h.Services.UserService.PostUser(r.Context(), user)
}

func toUser(userRequest PostUserRequest) (appUser.User, error) {
	hashedPassword, err := utils.HashPassword(userRequest.Password)
	if err != nil {
		return appUser.User{}, err
	}
	return appUser.User{
		Name:           userRequest.Name,
		Email:          userRequest.Email,
		HashedPassword: hashedPassword,
		Role:           userRequest.Role,
		CorporationId:  userRequest.CorporationId,
	}, nil
}

func (h *Handler) GetUser(c *gin.Context) {
}

func (h *Handler) UpdateUser(c *gin.Context) {
}

func (h *Handler) DeleteUser(c *gin.Context) {
}
