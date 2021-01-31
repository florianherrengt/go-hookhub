package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/florianherrengt/hubhook/config"
	"github.com/florianherrengt/hubhook/pkg/datasource"
	"github.com/florianherrengt/hubhook/pkg/worker"
	"github.com/jinzhu/configor"
)

func main() {
	err := configor.Load(&config.Config, "config.yml")
	if err != nil {
		log.Fatal(err)
	}
	datasource.ConnectStreamClient("db-worker")
	defer datasource.StreamClient.Close()
	qsub := worker.NewDBWorker()
	defer qsub.Close()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	log.Println()
	log.Printf("Draining...")
	datasource.StreamClient.NatsConn().Drain()
	log.Fatalf("Exiting")
}
