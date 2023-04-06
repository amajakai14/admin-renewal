package http

import (
	"fmt"
	"net/http"

	channel "github.com/amajakai14/admin-renewal/internal/channel"
	"github.com/amajakai14/admin-renewal/internal/course"
	"github.com/amajakai14/admin-renewal/internal/menu"
	appUser "github.com/amajakai14/admin-renewal/internal/user"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Router   *gin.Engine
	Services Services
	Server   *http.Server
}

type Services struct {
	UserService *appUser.Service
	ChannelService  *channel.Service
	MenuService *menu.Service
	CourseService *course.Service
}

func NewHandler(services Services) *Handler {
	h := &Handler{
		Services: services,
	}
	h.Router = gin.Default()
	h.mapRoutes()

	h.Server = &http.Server{
		Addr:    ":8080",
		Handler: h.Router,
	}

	return h
}

func (h *Handler) mapRoutes() {
	h.Router.GET("/ping", func(c *gin.Context) {
		fmt.Fprint(c.Writer, "pong")
	})
	h.Router.POST("api/v1/users", h.PostUser)
	h.Router.GET("api/v1/me", AuthMiddleware(),h.GetUser)
	h.Router.POST("api/v1/users/signin", h.SignIn)
	h.Router.POST("api/v1/image", h.UploadMenuImage)
}

func (h *Handler) Serve() error {
	if err := h.Server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
	return nil
}
