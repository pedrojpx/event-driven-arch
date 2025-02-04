package kafka

import (
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type Consumer struct {
	ConfigMap *ckafka.ConfigMap
	Topics    []string
}

func NewConsumer(configMap *ckafka.ConfigMap, topics []string) *Consumer {
	return &Consumer{
		ConfigMap: configMap,
		Topics:    topics,
	}
}

func (c *Consumer) Consume(msgChan chan *ckafka.Message) error {
	consumer, err := ckafka.NewConsumer(c.ConfigMap)
	if err != nil {
		panic(err)
	}
	// fmt.Println("created new consumer with")
	// fmt.Println(c.ConfigMap)

	err = consumer.SubscribeTopics(c.Topics, nil)
	if err != nil {
		panic(err)
	}
	// fmt.Println("subscribed to topics")
	// fmt.Println(c.Topics)

	for {
		// fmt.Println("waiting msg in consume function")
		msg, err := consumer.ReadMessage(-1)
		// fmt.Println("got it in consume function")
		if err == nil {
			// fmt.Println(string(msg.Value))
			msgChan <- msg
		} else {
			fmt.Println(err.Error())
		}
	}
}
