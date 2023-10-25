package handler

import (
	"net/http"

	"github.com/RenanAlmeida225/go-upload-images/internal/schemas"
	"github.com/RenanAlmeida225/go-upload-images/pkg/crypto"
	"github.com/RenanAlmeida225/go-upload-images/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func SignIn(ctx *gin.Context) {
	request := SignInRequest{}
	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := schemas.User{}
	if err := db.First(&user, "email=?", request.Email).Error; err != nil {
		sendError(ctx, http.StatusUnauthorized, "email invalid")
		return
	}

	if !user.IsEnable {
		sendError(ctx, http.StatusUnauthorized, "confirm your email")
		return
	}

	if err := crypto.ComparePasswords(user.Password, request.Password); err != nil {
		sendError(ctx, http.StatusUnauthorized, "password invalid")
		return
	}

	jwt, err := jwt.GenerateJwt(user.Email)

	if err != nil {
		sendError(ctx, http.StatusBadRequest, "error on created jwt")
		return
	}

	send(ctx, http.StatusOK, "login user", jwt)
}
