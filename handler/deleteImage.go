package handler

import (
	"net/http"

	"github.com/RenanAlmeida225/go-upload-images/helper"
	"github.com/RenanAlmeida225/go-upload-images/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteImage(ctx *gin.Context) {
	id := ctx.Param("id")

	image := schemas.Images{}

	if err := db.First(&image, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "image not found")
		return
	}

	if err := helper.DeleteInS3(image.Name); err != nil { // delete on aws
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := db.Delete(&image).Error; err != nil { // delete on database
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	send(ctx, http.StatusOK, "delete image", image)
}
