package handler

import (
	"net/http"

	"github.com/RenanAlmeida225/go-upload-images/internal/schemas"
	"github.com/gin-gonic/gin"
)

func GetImage(ctx *gin.Context) {
	id := ctx.Param("id")
	image := schemas.Images{}
	if err := db.First(&image, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "image not found")
		return
	}

	send(ctx, http.StatusOK, "get an image", image)
}
