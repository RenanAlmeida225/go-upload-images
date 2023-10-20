package handler

import (
	"github.com/RenanAlmeida225/go-upload-images/config"
	"github.com/RenanAlmeida225/go-upload-images/helper"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitializeHandler() {
	helper.InitializeHelper()
	db = config.GetPostgres()
}
