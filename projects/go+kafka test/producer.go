package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:56086",
	    "acks": "all",
	})
	if err != nil {
		fmt.Printf("failed to create producer: %s\n", err)
		os.Exit(1)
	}

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Produced event to topic %s: key = %-10s value = %s\n",
						*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value),
					)
				}
			}
		}
	}()

	users := [...]string{"eabara", "jsmith", "sgarcia", "jbernard", "htanaka", "awalther"}
    items := [...]string{"book", "alarm clock", "t-shirts", "gift card", "batteries"}
    topic := "purchases"

	for n := 0; n < 10; n++ {
		key := users[rand.Intn(len(users))]
		value := items[rand.Intn(len(items))]
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Key: []byte(key),
			Value: []byte(value),
		}, nil)
	}

	p.Flush(15 * 1000)
	p.Close()
}