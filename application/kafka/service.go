package kafka

import (
	"context"
	"fmt"

	"github.com/c-4u/employee-service/application/kafka/schema"
	"github.com/c-4u/employee-service/domain/service"
	"github.com/c-4u/employee-service/infrastructure/external"
	"github.com/c-4u/employee-service/infrastructure/external/topic"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaProcessor struct {
	Service *service.Service
	K       *external.Kafka
}

func NewKafkaProcessor(service *service.Service, kafka *external.Kafka) *KafkaProcessor {
	return &KafkaProcessor{
		Service: service,
		K:       kafka,
	}
}

func (p *KafkaProcessor) Consume() {
	p.K.Consumer.SubscribeTopics(p.K.ConsumeTopics, nil)
	for {
		msg, err := p.K.Consumer.ReadMessage(-1)
		if err == nil {
			// fmt.Println(string(msg.Value))
			p.processMessage(msg)
		}
	}
}

func (p *KafkaProcessor) processMessage(msg *ckafka.Message) {
	switch _topic := *msg.TopicPartition.Topic; _topic {
	case topic.NEW_USER:
		// TODO: add fault tolerance
		err := p.createUser(msg)
		if err != nil {
			fmt.Println("creation error ", err)
		}
	default:
		fmt.Println("not a valid topic", string(msg.Value))
	}
}

func (p *KafkaProcessor) createUser(msg *ckafka.Message) error {
	userEvent := schema.NewUserEvent()
	err := userEvent.ParseJson(msg.Value)
	if err != nil {
		return err
	}

	err = p.Service.CreateUser(context.TODO(), userEvent.User.ID, userEvent.User.Username, userEvent.User.EmployeeID)
	if err != nil {
		return err
	}

	return nil
}
