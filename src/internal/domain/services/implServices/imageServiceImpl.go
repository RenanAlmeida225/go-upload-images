package implServices

import (
	"fmt"
	"time"

	"github.com/RenanAlmeida225/go-upload-images/src/internal/api/dto"
	"github.com/RenanAlmeida225/go-upload-images/src/internal/domain/entities"
	"github.com/RenanAlmeida225/go-upload-images/src/internal/domain/repositories"
)

type ImageServiceImpl struct {
	repository repositories.ImageRepository
}

func New(repository repositories.ImageRepository) *ImageServiceImpl {
	return &ImageServiceImpl{repository: repository}
}

func (i *ImageServiceImpl) Save(image entities.Image) {
	dateTime := time.Now().Format("2006-01-02T15:04:05")
	image.Name = fmt.Sprintf("%s-%s", dateTime, image.Name)
	i.repository.Save(image)
}

func (i *ImageServiceImpl) GetImages(page, offset int) (imagesDto []dto.ImageDto) {
	images := i.repository.GetImages(page, offset)
	for _, image := range images {
		imagesDto = append(imagesDto, dto.ImageDto{ID: image.ID, Title: image.Title, Description: image.Description, Url: image.Url})
	}
	return imagesDto
}

func (i *ImageServiceImpl) GetImageById(id int) (dto.ImageDto, error) {
	image, err := i.repository.GetImageById(id)
	if err != nil {
		return dto.ImageDto{}, err
	}
	imageDto := dto.ImageDto{ID: image.ID, Title: image.Title, Description: image.Description, Url: image.Url}
	return imageDto, nil
}

func (i *ImageServiceImpl) GetImagesByTitleOrDescription(data string) (imagesDto []dto.ImageDto) {
	images := i.repository.GetImagesByTitleOrDescription(data)
	for _, image := range images {
		imagesDto = append(imagesDto, dto.ImageDto{ID: image.ID, Title: image.Title, Description: image.Description, Url: image.Url})
	}
	return imagesDto
}

func (i *ImageServiceImpl) UpdateImage(imageDto dto.ImageDto) error {
	err := i.repository.UpdateImage(imageDto)
	if err != nil {
		return err
	}
	return nil
}

func (i *ImageServiceImpl) DeleteImage(id int) error {
	err := i.repository.DeleteImage(id)
	if err != nil {
		return err
	}
	return nil
}
