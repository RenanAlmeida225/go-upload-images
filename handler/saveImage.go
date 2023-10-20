package handler

import (
	"net/http"
	"time"

	"github.com/RenanAlmeida225/go-upload-images/helper"
	"github.com/RenanAlmeida225/go-upload-images/schemas"
	"github.com/gin-gonic/gin"
)

func SaveImageHandler(ctx *gin.Context) {
	request := SaveImageRequest{}
	ctx.ShouldBind(&request)

	if err := request.Validate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	now := time.Now()
	name := now.Format("20060102150405") + "_" + request.Image.Filename

	url := helper.SaveInS3() // save on aws s3

	s := schemas.Images{
		Title:        request.Title,
		Description:  request.Description,
		Name:         name,
		OriginalName: request.Image.Filename,
		MimeType:     request.Image.Header.Get("Content-Type"),
		Url:          url,
	}
	// save in postgres
	send(ctx, http.StatusCreated, "save image", s)
}
