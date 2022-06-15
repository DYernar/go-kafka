package main

import (
	"log"
	"os"

	"github.com/Shopify/sarama"
)

func main() {
	// create kafka consumer on port 4444
	kafkaURL := os.Getenv("kafkaURL")
	topic := os.Getenv("topic")
	brokers := []string{kafkaURL}
	consumer, err := sarama.NewConsumer(brokers, nil)
	if err != nil {
		log.Printf("Error creating consumer: %v", err)
		return
	}

	partitionList, err := consumer.Partitions(topic) //get all partitions
	if err != nil {
		log.Printf("Error getting partition list: %v", err)
		return
	}

	log.Printf("Connected to kafka at %s", kafkaURL)
	log.Printf("Partitions: %v", partitionList)

	initialOffset := sarama.OffsetOldest //offset to start reading message from
	pc, _ := consumer.ConsumePartition(topic, partitionList[0], initialOffset)
	for message := range pc.Messages() {
		log.Printf("Message received: %s", message.Value)
	}
	log.Printf("Finishing consuming")
}
