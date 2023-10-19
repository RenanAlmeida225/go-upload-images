package router

import (
	"github.com/RenanAlmeida225/go-upload-images/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	v1.POST("/images", handler.SaveImageHandler)
}
