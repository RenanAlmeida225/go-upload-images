package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	v1.POST("/images", func(ctx *gin.Context) {
		ctx.JSON(http.StatusCreated, gin.H{
			"msg": "created",
		})
	})
}
