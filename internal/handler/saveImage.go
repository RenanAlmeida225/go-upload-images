package handler

import (
	"net/http"
	"time"

	"github.com/RenanAlmeida225/go-upload-images/internal/schemas"
	"github.com/RenanAlmeida225/go-upload-images/pkg/s3"
	"github.com/gin-gonic/gin"
)

func SaveImageHandler(ctx *gin.Context) {
	email, _ := ctx.Get("email")
	user := schemas.User{}
	if err := db.First(&user, "email=?", email).Error; err != nil {
		sendError(ctx, http.StatusUnauthorized, "user not found")
		return
	}

	request := SaveImageRequest{}
	ctx.ShouldBind(&request)

	if err := request.Validate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	now := time.Now()
	name := now.Format("20060102150405") + "_" + request.Image.Filename

	url, err := s3.SaveInS3(request.Image, name) // save on aws s3
	if err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	image := schemas.Images{
		Title:        request.Title,
		Description:  request.Description,
		Name:         name,
		OriginalName: request.Image.Filename,
		MimeType:     request.Image.Header.Get("Content-Type"),
		Url:          url,
		UserId:       user.ID,
	}

	user.Images = append(user.Images, image)

	if err := db.Save(&user).Error; err != nil { // save in database
		sendError(ctx, http.StatusBadRequest, err.Error())
	}

	send(ctx, http.StatusCreated, "save image", &image)
}
