package main

import (
	"fmt"

	"github.com/RenanAlmeida225/go-upload-images/config"
	"github.com/RenanAlmeida225/go-upload-images/router"
)

func main() {
	err := config.Init()
	if err != nil {
		fmt.Println(err)
	}
	router.Initialize()
}
