package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Shopify/sarama"
)

func main() {
	kafkaURL := os.Getenv("kafkaURL")
	topic := os.Getenv("topic")
	brokers := []string{kafkaURL}
	//setup relevant config info
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	log.Printf("Connecting to Kafka at %s", kafkaURL)
	producer, err := sarama.NewSyncProducer(brokers, config)

	if err != nil {
		log.Printf("Error while creating new producer: %v", err)
		return
	}

	log.Printf("Connected to Kafka at %s", kafkaURL)

	for i := 0; i < 10000; i++ {
		var partition int32 = 0                                         //Partition to produce to
		msg := fmt.Sprintf("actual information to save on kafka %d", i) //e.g {"name":"John Doe", "email":"john.doe@email.com"}
		message := &sarama.ProducerMessage{
			Topic:     topic,
			Partition: partition,
			Value:     sarama.StringEncoder(msg),
		}
		partition, offset, err := producer.SendMessage(message)
		if err != nil {
			log.Printf("Error while producing message: %v", err)
			return
		}
		log.Printf("Message '%s' published to topic '%s' partition %d at offset %d %d", msg, topic, partition, offset, i)
		time.Sleep(500 * time.Millisecond)
	}
}
