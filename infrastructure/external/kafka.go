package external

import (
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type Kafka struct {
	Consumer      *ckafka.Consumer
	Producer      *ckafka.Producer
	ConsumeTopics []string
	DeliveryChan  chan ckafka.Event
}

func NewKafka(servers string, groupId string, consumeTopics []string, deliveryChan chan ckafka.Event) (*Kafka, error) {
	p, err := ckafka.NewProducer(
		&ckafka.ConfigMap{
			"bootstrap.servers": servers,
		},
	)
	if err != nil {
		return nil, err
	}

	c, err := ckafka.NewConsumer(
		&ckafka.ConfigMap{
			"bootstrap.servers": servers,
			"group.id":          groupId,
			"auto.offset.reset": "earliest",
		},
	)
	if err != nil {
		return nil, err
	}

	return &Kafka{
		Consumer:      c,
		Producer:      p,
		ConsumeTopics: consumeTopics,
		DeliveryChan:  deliveryChan,
	}, nil
}

// TODO: Add event log
func (k *Kafka) DeliveryReport() {
	for e := range k.DeliveryChan {
		switch ev := e.(type) {
		case *ckafka.Message:
			if ev.TopicPartition.Error != nil {
				// TODO: add attempts
				fmt.Println("Delivery failed:", ev.TopicPartition)
			} else {
				fmt.Println("Delivered message to:", ev.TopicPartition)
			}
		}
	}
}
