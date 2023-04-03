package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func main() {
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: "sing",
		Value: sarama.StringEncoder("testing 123"),
	}

	p, o, err := producer.SendMessage(msg)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Partition: %d, Offset: %d", p, o)
}
