package main

import (
	"fmt"

	"github.com/IBM/sarama"
)

func main() {
	servers := []string{"localhost:9092"}

	consumer, err := sarama.NewConsumer(servers, nil)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("king", 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}
	defer partitionConsumer.Close()
	fmt.Println("Consumer Start..")
	for {
		select {
		case err := <-partitionConsumer.Errors():
			fmt.Println(err)
		case msg := <-partitionConsumer.Messages():
			fmt.Println(msg.Topic)
			fmt.Println(string(msg.Value))

		}
	}

}
