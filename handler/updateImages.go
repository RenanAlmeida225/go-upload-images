package handler

import (
	"net/http"

	"github.com/RenanAlmeida225/go-upload-images/schemas"
	"github.com/gin-gonic/gin"
)

func UpateImage(ctx *gin.Context) {
	id := ctx.Param("id")
	request := UpdateImageRequest{}
	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	image := schemas.Images{}

	if err := db.First(&image, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "image not found")
		return
	}

	// update values
	if request.Title != "" {
		image.Title = request.Title
	}
	if request.Description != "" {
		image.Description = request.Description
	}

	if err := db.Save(&image).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error updating image")
		return
	}

	send(ctx, http.StatusOK, "uptade image", image)
}
