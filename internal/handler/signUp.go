package handler

import (
	"fmt"
	"net/http"

	"github.com/RenanAlmeida225/go-upload-images/internal/helper"
	"github.com/RenanAlmeida225/go-upload-images/internal/schemas"
	"github.com/RenanAlmeida225/go-upload-images/pkg/crypto"
	"github.com/RenanAlmeida225/go-upload-images/pkg/uuid"
	"github.com/gin-gonic/gin"
)

func SignUp(ctx *gin.Context) {
	request := SignUpRequest{}
	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := schemas.User{}
	if err := db.First(&user, "email=?", request.Email).Error; err == nil {
		sendError(ctx, http.StatusBadRequest, "email invalid")
		return
	}

	passwordHash, err := crypto.GeneratePassword(request.Password)
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "Server error")
		return
	}

	user.Username = request.Username
	user.Email = request.Email
	user.Password = passwordHash

	token := uuid.GenerateConfirmationToken()

	confirmation := schemas.Confirmation{
		Token:  token,
		UserID: user.ID,
		User:   user,
	}

	if err := db.Create(&confirmation).Error; err != nil {
		sendError(ctx, http.StatusBadRequest, "fail on register user")
		return
	}

	if err := helper.SendEmailConfimation(user.Username, user.Email, token); err != nil {
		fmt.Println(err)
		sendError(ctx, http.StatusBadRequest, "fail on send email")
		return
	}

	send(ctx, http.StatusCreated, "sign up", "register user successfully")

}
