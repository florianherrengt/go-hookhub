package main

import (
	"log"
	"os"

	"github.com/florianherrengt/hubhook/config"
	"github.com/florianherrengt/hubhook/pkg/datasource"
	"github.com/florianherrengt/hubhook/pkg/routers"
	"github.com/jinzhu/configor"
)

func main() {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "config.yml"
	}
	log.Println("using config file from ", configPath)
	err := configor.Load(&config.Config, configPath)
	if err != nil {
		log.Fatal(err)
	}
	datasource.ConnectStreamClient("api")
	datasource.ConnectDataBase()
	r := routers.NewRouter()
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	r.Run(":" + port)
}
