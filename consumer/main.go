package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func main() {
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		panic(err)
	}

	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("sing", 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}

	defer partitionConsumer.Close()


	fmt.Println("Waiting for messages. To exit press CTRL+C")
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			fmt.Println(string(msg.Value))
		case err := <-partitionConsumer.Errors():
			fmt.Println(err.Error())
		}
	}
}
