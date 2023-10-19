package handler

import (
	"net/http"
	"strings"
	"time"

	"github.com/RenanAlmeida225/go-upload-images/helper"
	"github.com/RenanAlmeida225/go-upload-images/schemas"
	"github.com/gin-gonic/gin"
)

func SaveImageHandler(ctx *gin.Context) {
	image, err := ctx.FormFile("image")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "fail on save image",
		})
		return
	}

	if !strings.Contains(image.Header.Get("Content-Type"), "image/") {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "fail on save image",
		})
		return
	}

	title := ctx.PostForm("title")
	description := ctx.PostForm("description")

	now := time.Now()
	name := now.Format("20060102150405") + "_" + image.Filename

	url := helper.SaveInS3()

	s := schemas.Images{
		Title:        title,
		Description:  description,
		Name:         name,
		OriginalName: image.Filename,
		MimeType:     image.Header.Get("Content-Type"),
		Url:          url,
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": s,
	})
}
