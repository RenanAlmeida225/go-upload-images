package repositories

import (
	"github.com/RenanAlmeida225/go-upload-images/src/internal/api/dto"
	"github.com/RenanAlmeida225/go-upload-images/src/internal/domain/entities"
)

type ImageRepository interface {
	Save(image entities.Image)
	GetImages(page, offset int) []entities.Image
	GetImageById(id int) (entities.Image, error)
	GetImagesByTitleOrDescription(data string) []entities.Image
	UpdateImage(imageDto dto.ImageDto) error
	DeleteImage(id int) error
}
