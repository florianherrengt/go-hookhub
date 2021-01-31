package datasource

import (
	"fmt"
	"log"

	"github.com/nats-io/stan.go"
)

// StreamClient is the stan connection to NATS
var StreamClient stan.Conn

// ConnectStreamClient is a function to create a new streaming client for NATS
func ConnectStreamClient(clientID string) {

	// var opts []stan.Option

	sc, err := stan.Connect("test-cluster", clientID)
	fmt.Println("connected")
	if err != nil {
		log.Panicln(err)
	}

	StreamClient = sc
}
