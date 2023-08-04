package impl

import (
	"errors"
	"log"

	"github.com/RenanAlmeida225/go-upload-images/src/internal/api/dto"
	"github.com/RenanAlmeida225/go-upload-images/src/internal/domain/entities"
	"gorm.io/gorm"
)

type ImageRepositoryImpl struct {
	DB *gorm.DB
}

func New(DB *gorm.DB) *ImageRepositoryImpl {
	return &ImageRepositoryImpl{DB: DB}
}

func (i *ImageRepositoryImpl) Save(image entities.Image) {
	if result := i.DB.Create(&image); result.Error != nil {
		log.Fatal("Repository: ", result.Error)
	}
}

func (i *ImageRepositoryImpl) GetImages(page, offset int) (images []entities.Image) {
	i.DB.Find(&images)
	return images
}

func (i *ImageRepositoryImpl) GetImageById(id int) (image entities.Image, err error) {
	if result := i.DB.First(&image, id); result.Error != nil {
		return entities.Image{}, errors.New("image not found")
	}
	return image, nil
}

func (i *ImageRepositoryImpl) GetImagesByTitleOrDescription(data string) (images []entities.Image) {
	i.DB.Find(&images)
	return images
}

func (i *ImageRepositoryImpl) UpdateImage(imageDto dto.ImageDto) error {
	var image entities.Image
	if result := i.DB.First(&image, imageDto.ID); result.Error != nil {
		return errors.New("image not found")
	}
	image.Title = imageDto.Title
	image.Description = imageDto.Description
	i.DB.Save(&image)
	return nil
}

func (i *ImageRepositoryImpl) DeleteImage(id int) error {
	var image entities.Image
	if result := i.DB.First(&image, id); result.Error != nil {
		return errors.New("image not found")
	}
	i.DB.Delete(&image)
	return nil
}
