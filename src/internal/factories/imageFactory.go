package factories

import (
	"github.com/RenanAlmeida225/go-upload-images/src/internal/api/controllers"
	"github.com/RenanAlmeida225/go-upload-images/src/internal/domain/repositories/impl"
	"github.com/RenanAlmeida225/go-upload-images/src/internal/domain/services/implServices"
	"gorm.io/gorm"
)

func ImageControllerFactory(service *implServices.ImageServiceImpl) *controllers.ImageController {
	return controllers.New(service)
}

func ImageServiceFactory(repository *impl.ImageRepositoryImpl) *implServices.ImageServiceImpl {
	return implServices.New(repository)
}

func ImageRepositoryFactory(DB *gorm.DB) *impl.ImageRepositoryImpl {
	return impl.New(DB)
}
