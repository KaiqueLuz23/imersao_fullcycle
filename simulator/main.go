package main

import (
	"fmt"

	"github.com/kaiqueluz/imersapfsfc2-simulator/infra/kafka"
)

func main() {
	msgChan := make(chan *ckafka.Message)
	cosumer := kafka.NewKafkaConsumer(msgChan)
	go cosumer.Consume()

	for msg := range msgChan {
		fmt.Println(string(msg.Value))
	}
	//producer := kafka.NewKafkaProducer()
	//kafka.Publish("Olá", "readtest", producer)
	// lop infinito
	//for {
	//	_ = 1
	//}
	//route := route.Route{
	//	ID:       "2",
	//	ClientID: "2",
	//}
	//route.LoadPositions()
	//stringjson, _ := route.ExportJsonPositions()
	//fmt.Println(stringjson[1])
}
