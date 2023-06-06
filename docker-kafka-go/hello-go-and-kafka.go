package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	kafkaBrokers    = "localhost:9092" // Comma-separated list of Kafka brokers
	kafkaTopic      = "telenor_topic"       // Kafka topic to produce and consume events
	consumerGroupID = "my_group"       // Consumer group ID for the Kafka consumer
)

func main() {
	// Create a new Kafka consumer configuration
	consumerConfig := &kafka.ConfigMap{
		"bootstrap.servers":  kafkaBrokers,
		"group.id":           consumerGroupID,
		"auto.offset.reset":  "earliest",
		"enable.auto.commit": "false",
	}

	// Create a new Kafka consumer
	consumer, err := kafka.NewConsumer(consumerConfig)
	if err != nil {
		log.Fatal("Failed to create Kafka consumer:", err)
	}
	defer consumer.Close()

	// Subscribe to the Kafka topic
	err = consumer.SubscribeTopics([]string{kafkaTopic}, nil)
	if err != nil {
		log.Fatal("Failed to subscribe to Kafka topic:", err)
	}

	// Create a new Kafka producer configuration
	producerConfig := &kafka.ConfigMap{
		"bootstrap.servers": kafkaBrokers,
	}

	// Create a new Kafka producer
	producer, err := kafka.NewProducer(producerConfig)
	if err != nil {
		log.Fatal("Failed to create Kafka producer:", err)
	}
	defer producer.Close()

	// Create a channel to handle OS interrupts
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// Create a wait group to wait for all goroutines to finish
	wg := &sync.WaitGroup{}
	wg.Add(2) // One for the producer and one for the consumer

	// Start the Kafka producer
	go runProducer(producer, wg)

	// Start the Kafka consumer
	go runConsumer(consumer, wg)

	// Wait for an interrupt signal
	<-signals

	// Close the Kafka consumer
	consumer.Close()

	// Wait for all goroutines to finish
	wg.Wait()
}

// runProducer generates and sends events to the Kafka topic
func runProducer(producer *kafka.Producer, wg *sync.WaitGroup) {
	defer wg.Done()

	// Generate and send events to the Kafka topic
	for i := 0; i < 10; i++ {
		message := fmt.Sprintf("Event %d", i)
		topic := kafkaTopic // Use a string variable instead of a constant
		err := producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(message),
		}, nil)
		if err != nil {
			log.Println("Failed to send message to Kafka:", err)
		} else {
			log.Printf("Message sent: %s\n", message)
		}
	}

	producer.Flush(1000)
}

// runConsumer consumes events from the Kafka topic
func runConsumer(consumer *kafka.Consumer, wg *sync.WaitGroup) {
	defer wg.Done()

	// Consume events from the Kafka topic
	for {
		ev := consumer.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			if e.TopicPartition.Error != nil {
				log.Println("Error consuming message from Kafka:", e.TopicPartition.Error)
				continue
			}

			log.Printf("Received message: Partition %d, Offset %d, Value: %s\n",
				e.TopicPartition.Partition, e.TopicPartition.Offset, string(e.Value))
		case kafka.Error:
			log.Println("Error consuming messages from Kafka:", e)
		}
	}
}
