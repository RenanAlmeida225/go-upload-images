package router

import (
	"github.com/RenanAlmeida225/go-upload-images/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {
	handler.InitializeHandler()
	v1 := r.Group("/api/v1")
	{
		v1.POST("/images", handler.SaveImageHandler)
		v1.GET("/images", handler.GetAllImages)
		v1.GET("/images/:id", handler.GetImage)
		v1.PATCH("/images/:id", handler.UpateImage)
		v1.DELETE("/images/:id", handler.DeleteImage)
	}

}
