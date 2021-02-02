package worker

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/florianherrengt/hubhook/config"
	"github.com/florianherrengt/hubhook/pkg/datasource"
	"github.com/florianherrengt/hubhook/pkg/models"
	"github.com/nats-io/stan.go"
)

// NewDBWorker is creating a new worker to
func NewDBWorker() stan.Subscription {
	fmt.Println("starting db worker")
	aw, _ := time.ParseDuration("10s")
	qsub, err := datasource.StreamClient.QueueSubscribe(config.Config.PubSub.EventName.NewIncomingRequest,
		"db-worker", func(m *stan.Msg) {
			hookEvent := models.HookEvent{}
			err := json.Unmarshal(m.Data, &hookEvent)
			if err != nil {
				log.Fatalln(err)
				return
			}
			result := datasource.DB.Create(&hookEvent)
			if result.Error != nil {
				log.Fatal(result.Error)
				return
			}
			err = m.Ack()
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
