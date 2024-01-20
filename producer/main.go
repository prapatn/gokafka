package main

import "github.com/IBM/sarama"

func main() {
	servers := []string{"localhost:9092"}

	producer, err := sarama.NewAsyncProducer(servers, nil)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	msg := sarama.ProducerMessage{
		Topic: "king",
		Value: sarama.StringEncoder("Hello King"),
	}

	envelopes := make(chan *sarama.ProducerMessage, 256)
	envelopes <- &msg
	select {
	case envelope := <-envelopes:
		producer.Input() <- envelope
	}
}
