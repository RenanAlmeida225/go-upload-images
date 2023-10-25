package handler

import (
	"github.com/RenanAlmeida225/go-upload-images/config"
	"github.com/RenanAlmeida225/go-upload-images/pkg"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitializeHandler() {
	pkg.InitializePkg()
	db = config.GetPostgres()
}
