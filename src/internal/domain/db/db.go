package db

import (
	"log"

	"github.com/RenanAlmeida225/go-upload-images/src/internal/domain/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbUrl := "postgres://postgres:p@ssw0rd@localhost:5432/postgres"
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&entities.Image{})

	return db
}
