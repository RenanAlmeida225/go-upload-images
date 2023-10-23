package router

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20
	initializeRoutes(r)
	log.Fatal(r.Run(":3000"))
}
