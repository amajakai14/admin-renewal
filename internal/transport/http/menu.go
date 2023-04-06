package http

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
)

type PostMenuRequest struct {
	MenuNameTH string `json:"menu_name_th"`
	MenuNameEN string `json:"menu_name_en"`
	MenuType   string `json:"menu_type"`
	Price      uint32 `json:"price"`
}

func (h *Handler) PostMenu(c *gin.Context) {

}

func (h *Handler) UploadMenuImage(c *gin.Context) {
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 10<<20)
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("error read body")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	img, err := jpeg.Decode(bytes.NewReader(body))
	if err != nil {
		fmt.Println("error decode image")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	output := resize.Resize(300, 300, img, resize.Lanczos3)

	c.Header("Content-Type", "image/jpeg")
	c.Writer.WriteHeader(http.StatusOK)
	if err := jpeg.Encode(c.Writer, output, &jpeg.Options{Quality: 100}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
}
