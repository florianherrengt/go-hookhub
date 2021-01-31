package worker

import (
	"fmt"
	"log"
	"time"

	"github.com/florianherrengt/hubhook/config"
	"github.com/florianherrengt/hubhook/pkg/datasource"
	"github.com/nats-io/stan.go"
)

// NewDBWorker is creating a new worker to
func NewDBWorker() stan.Subscription {
	fmt.Println("starting db worker")
	aw, _ := time.ParseDuration("10s")
	qsub, err := datasource.StreamClient.QueueSubscribe(config.Config.PubSub.EventName.NewIncomingRequest,
		"db-worker", func(m *stan.Msg) {
			fmt.Println("1", string(m.Data))
			err := m.Ack()
			if err != nil {
				log.Fatalln(err)
			}
		}, stan.DurableName("db-worker"),
		stan.DeliverAllAvailable(), stan.SetManualAckMode(), stan.AckWait(aw))

	if err != nil {
		log.Panicln(err)
	}

	return qsub
}
