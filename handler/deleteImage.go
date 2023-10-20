package handler

import (
	"net/http"

	"github.com/RenanAlmeida225/go-upload-images/helper"
	"github.com/RenanAlmeida225/go-upload-images/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteImage(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	image := schemas.Images{}

	if err := db.First(&image, id).Error; err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := helper.DeleteInS3(image.Name); err != nil { // delete on aws
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := db.Delete(&image).Error; err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	send(ctx, http.StatusOK, "delete image", image)
}
