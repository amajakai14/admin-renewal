package http

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type PostChannelRequest struct {
	TableID  int `json:"table_id" validate:"required"`
	CourseID int `json:"course_id" validate:"required"`
}

func (h *Handler) PostChannel(c *gin.Context) {
	r := c.Request
	w := c.Writer
	var channelRequest PostChannelRequest

	if err := json.NewDecoder(r.Body).Decode(&channelRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	validate := validator.New()
	err := validate.Struct(channelRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	//TODO add service h.Services.ChannelService.PostChannel(c, channelRequest)
}
