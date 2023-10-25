package handler

import (
	"net/http"

	"github.com/RenanAlmeida225/go-upload-images/internal/schemas"
	"github.com/gin-gonic/gin"
)

func ConfirmationEmail(ctx *gin.Context) {
	token := ctx.Param("token")

	confirmation := schemas.Confirmation{}
	user := &confirmation.User

	if err := db.Where("token=?", token).Preload("User").First(&confirmation).Error; err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user.IsEnable = true

	if err := db.Save(&user).Error; err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := db.Delete(&confirmation).Error; err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.Status(http.StatusNoContent)
}
