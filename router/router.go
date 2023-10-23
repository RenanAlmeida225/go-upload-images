package router

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20
	initializeRoutes(r)
	log.Fatal(r.Run(":3000"))
}
