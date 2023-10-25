package handler

import (
	"net/http"
	"strconv"

	"github.com/RenanAlmeida225/go-upload-images/internal/schemas"
	"github.com/gin-gonic/gin"
)

func GetAllImages(ctx *gin.Context) {
	email, _ := ctx.Get("email")
	user := schemas.User{}
	if err := db.First(&user, "email=?", email).Error; err != nil {
		sendError(ctx, http.StatusUnauthorized, "user not found")
		return
	}

	page, _ := strconv.Atoi(ctx.Query("page"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	offset := (page - 1) * limit

	images := []schemas.Images{}

	if err := db.Offset(offset).Limit(limit).Find(&images, "user_id=?", user.ID).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "images not found")
		return
	}

	send(ctx, http.StatusOK, "get all images", images)
}
