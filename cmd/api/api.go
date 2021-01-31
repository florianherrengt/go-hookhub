package main

import (
	"log"
	"os"

	"github.com/florianherrengt/hubhook/config"
	"github.com/florianherrengt/hubhook/pkg/router"
	"github.com/jinzhu/configor"
)

func main() {
	err := configor.Load(&config.Config, "config.yml")
	if err != nil {
		log.Fatal(err)
	}
	// datasource.ConnectStreamClient("api")
	r := router.NewRouter()
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	r.Run(":" + port)
}
