package services

import (
	"github.com/RenanAlmeida225/go-upload-images/src/internal/api/dto"
	"github.com/RenanAlmeida225/go-upload-images/src/internal/domain/entities"
)

type ImageService interface {
	Save(image entities.Image)
	GetImages(page, offset int) []dto.ImageDto
	GetImageById(id int) (dto.ImageDto, error)
	GetImagesByTitleOrDescription(data string) []dto.ImageDto
	UpdateImage(imageDto dto.ImageDto) error
	DeleteImage(id int) error
}
