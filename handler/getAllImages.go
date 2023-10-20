package handler

import (
	"net/http"
	"strconv"

	"github.com/RenanAlmeida225/go-upload-images/schemas"
	"github.com/gin-gonic/gin"
)

func GetAllImages(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	offset := (page - 1) * limit

	images := []schemas.Images{}

	if err := db.Offset(offset).Limit(limit).Find(&images).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "images not found")
		return
	}

	send(ctx, http.StatusOK, "get all images", images)
}
