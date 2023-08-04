package controllers

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/RenanAlmeida225/go-upload-images/src/internal/api/dto"
	"github.com/RenanAlmeida225/go-upload-images/src/internal/domain/entities"
	"github.com/RenanAlmeida225/go-upload-images/src/internal/domain/services"
	"github.com/gin-gonic/gin"
)

type ImageController struct {
	service services.ImageService
}

func New(service services.ImageService) *ImageController {
	return &ImageController{service: service}
}

func (i *ImageController) Save(ctx *gin.Context) {
	title := ctx.PostForm("title")
	description := ctx.PostForm("description")
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	images := entities.Image{Title: title, Description: description,
		Name:         file.Filename,
		OriginalName: file.Filename,
		MimeType:     file.Header["Content-Type"][0],
		Url:          "url/file.png"}
	i.service.Save(images)
	ctx.JSON(http.StatusCreated, fmt.Sprintf("%s uploaded!", file.Filename))
}

func (i *ImageController) GetImages(ctx *gin.Context) {
	page := ctx.Query("page")
	offset := ctx.Query("offset")
	pageInt, _ := strconv.Atoi(page)
	offsetInt, _ := strconv.Atoi(offset)
	images := i.service.GetImages(pageInt, offsetInt)
	ctx.JSON(http.StatusOK, images)
}

func (i *ImageController) GetImageById(ctx *gin.Context) {
	id := ctx.Param("id")
	idToInt, _ := strconv.Atoi(id)
	image, err := i.service.GetImageById(idToInt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, image)
}
func (i *ImageController) UpdateImage(ctx *gin.Context) {
	id := ctx.Param("id")
	idToInt, _ := strconv.Atoi(id)
	b, _ := io.ReadAll(ctx.Request.Body)
	reqBody := string(b)
	regex := regexp.MustCompile(`[\{-\}]|\t|\"title\"|\"description\"|\:+\s|\"|\n`)
	title := strings.Split(regex.ReplaceAllString(reqBody, ""), ",")[0]
	description := strings.Split(regex.ReplaceAllString(reqBody, ""), ",")[1]
	err := i.service.UpdateImage(dto.ImageDto{ID: idToInt, Title: title, Description: description})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}

func (i *ImageController) DeleteImage(ctx *gin.Context) {
	id := ctx.Param("id")
	idToInt, _ := strconv.Atoi(id)
	err := i.service.DeleteImage(idToInt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}
