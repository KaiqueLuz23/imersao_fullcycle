package kafka

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/kaiqueluz/imersapfsfc2-simulator/application/route"
	"github.com/kaiqueluz/imersapfsfc2-simulator/infra/kafka"
)

func producer(msg *ckafka.Message) {
	producer := kafka.NewKafkaProducer()
	route := route.NewRoute()
	json.Unmarshal(msg.Value, &route)
	route.LoadPositions()
	positions, err := route.ExportJsonPositions()
	if err != nil {
		log.Println((err.Error()))
	}
	for _, p := range positions {
		kafka.Publish(p, os.Getenv("KafkaProduceTopic"), producer)
		time.Sleep(time.Millisecond * 500)
	}

}
