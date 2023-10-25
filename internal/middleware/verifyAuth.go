package middleware

import (
	"net/http"
	"strings"

	"github.com/RenanAlmeida225/go-upload-images/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func VerifyAuth(ctx *gin.Context) {
	header := ctx.GetHeader("Authorization")

	if header == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "bad header value given",
		})
		return
	}

	jwtString := strings.Split(header, " ")
	if len(jwtString) != 2 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "incorrectly formatted authorization header",
		})
		return
	}

	email, err := jwt.VerifyJwt(jwtString[1])
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.Set("email", email)
	ctx.Next()
}
