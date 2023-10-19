package handler

import (
	"net/http"
	"time"

	"github.com/RenanAlmeida225/go-upload-images/helper"
	"github.com/RenanAlmeida225/go-upload-images/schemas"
	"github.com/gin-gonic/gin"
)

func SaveImageHandler(ctx *gin.Context) {
	image, _ := ctx.FormFile("image")
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
