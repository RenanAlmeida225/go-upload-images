package router

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	r := gin.Default()
	initializeRoutes(r)
	log.Fatal(r.Run())
}
