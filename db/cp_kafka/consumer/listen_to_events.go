package consumer

import (
	"app-container-platform/config"
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func ListenEvents() {
	config.Log.Debug().Msg(fmt.Sprintf("Started listening to events to Broker: %s -- Group: %s -- topic: %s", config.KafkaBroker, config.KafkaConsumerGroup, config.KafkaListenerTopic))

	//this := new(listenEventsHelper)
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{config.KafkaBroker},
		GroupID:  config.KafkaConsumerGroup,
		Topic:    config.KafkaListenerTopic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	for {
		//var data _type.KafkaMessage
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Println(err.Error())
			continue
		}

		inputMessage := fmt.Sprintf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))

		log.Println(inputMessage)

		time.Sleep(3 * time.Second)
	}
}
