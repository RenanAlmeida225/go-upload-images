package handler

import (
	"net/http"

	"github.com/RenanAlmeida225/go-upload-images/internal/schemas"
	"github.com/gin-gonic/gin"
)

func GetImage(ctx *gin.Context) {
	email, _ := ctx.Get("email")
	user := schemas.User{}

	if err := db.First(&user, "email=?", email).Error; err != nil {
		sendError(ctx, http.StatusUnauthorized, "user not found")
		return
	}

	id := ctx.Param("id")
	image := schemas.Images{}
	if err := db.First(&image, "id=? and user_id=?", id, user.ID).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "image not found")
		return
	}

	send(ctx, http.StatusOK, "get an image", image)
}
