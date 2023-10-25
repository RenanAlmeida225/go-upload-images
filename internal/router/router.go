package router

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	gin.SetMode(os.Getenv("GIN_MODE"))
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20
	initializeRoutes(r)
	log.Fatal(r.Run(":" + os.Getenv("PORT")))
}
