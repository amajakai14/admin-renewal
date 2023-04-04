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
	PostUser(ctx context.Context, user *appUser.User) error
	GetUser(ctx context.Context, id int) (*appUser.User, error)
	UpdateUser(ctx context.Context, user *appUser.User) error
	DeleteUser(ctx context.Context, id int) error
}

type PostUserRequest struct {
	Name          string `json:"name"`
	Email         string `json:"email" validate:"required,email"`
	Password      string `json:"password" validate:"required,min=8"`
	Role          string `json:"role" validate:"required,oneof=admin staff"`
	CorporationId string `json:"corporation_id"`
}

type SignInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func (h *Handler) SignIn(c *gin.Context) {
	r := c.Request
	w := c.Writer
	var signInRequest SignInRequest 

	if err := json.NewDecoder(r.Body).Decode(&signInRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	validate := validator.New()
	err := validate.Struct(signInRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	user, err := h.Services.UserService.GetUserByEmail(r.Context(), signInRequest.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	match := utils.MatchPassword(signInRequest.Password, user.HashedPassword)
	if !match {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	token, err := Generate(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
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

type UserResponse struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Role          string `json:"role"`
	CorporationId string `json:"corporation_id"`
}

func (h *Handler) GetUser(c *gin.Context) {
	userID, exist := c.Get("userId")
	if exist == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user id not found",
		})
		return
	}
	id := int(userID.(float64))
	user, err := h.Services.UserService.GetUserByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, toUserResponse(&user))
}

func toUserResponse(user *appUser.User) UserResponse {
	return UserResponse{
		ID:            user.ID,
		Name:          user.Name,
		Email:         user.Email,
		Role:          user.Role,
		CorporationId: user.CorporationId,
	}
}

func (h *Handler) UpdateUser(c *gin.Context) {
}

func (h *Handler) DeleteUser(c *gin.Context) {
}
