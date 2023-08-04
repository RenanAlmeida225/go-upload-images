package main

import (
	"github.com/RenanAlmeida225/go-upload-images/src/internal/api/routes"
	"github.com/RenanAlmeida225/go-upload-images/src/internal/domain/db"

	"github.com/gin-gonic/gin"
)

func main() {
	DB := db.Init()
	router := gin.Default()
	routes.Routes(router, DB)
	router.Run()
}
