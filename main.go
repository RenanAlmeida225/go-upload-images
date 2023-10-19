package main

import (
	"log"

	"github.com/RenanAlmeida225/go-upload-images/config"
	"github.com/RenanAlmeida225/go-upload-images/router"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Fatal(err)
	}
	router.Initialize()
}
