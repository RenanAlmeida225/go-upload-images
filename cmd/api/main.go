package main

import (
	"github.com/RenanAlmeida225/go-upload-images/config"
	"github.com/RenanAlmeida225/go-upload-images/internal/router"
)

func main() {
	err := config.Init()
	if err != nil {
		panic(err)
	}
	router.Initialize()
}
