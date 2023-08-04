package routes

import (
	"github.com/RenanAlmeida225/go-upload-images/src/internal/factories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(router *gin.Engine, DB *gorm.DB) {
	imageRepository := factories.ImageRepositoryFactory(DB)
	imageService := factories.ImageServiceFactory(imageRepository)
	imageController := factories.ImageControllerFactory(imageService)
	router.POST("/images", imageController.Save)
	router.GET("/images", imageController.GetImages)
	router.GET("/images/:id", imageController.GetImageById)
	router.PUT("/images/:id", imageController.UpdateImage)
	router.DELETE("/images/:id", imageController.DeleteImage)
}
