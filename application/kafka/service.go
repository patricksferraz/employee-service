package kafka

import (
	"context"
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/patricksferraz/employee-service/application/kafka/schema"
	"github.com/patricksferraz/employee-service/domain/service"
	"github.com/patricksferraz/employee-service/infrastructure/external"
	"github.com/patricksferraz/employee-service/infrastructure/external/topic"
)

type KafkaProcessor struct {
	Service *service.Service
	K       *external.KafkaConsumer
}

func NewKafkaProcessor(service *service.Service, kafkaConsumer *external.KafkaConsumer) *KafkaProcessor {
	return &KafkaProcessor{
		Service: service,
		K:       kafkaConsumer,
	}
}

func (p *KafkaProcessor) Consume() {
	p.K.Consumer.SubscribeTopics(p.K.ConsumerTopics, nil)
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
	case topic.NEW_COMPANY:
		err := p.createCompany(msg)
		if err != nil {
			fmt.Println("creation error ", err)
		}
	case topic.ADD_EMPLOYEE_TO_COMPANY:
		err := p.addEmployeeToCompany(msg)
		if err != nil {
			fmt.Println("addition error ", err)
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

func (p *KafkaProcessor) createCompany(msg *ckafka.Message) error {
	companyEvent := schema.NewCompanyEvent()
	err := companyEvent.ParseJson(msg.Value)
	if err != nil {
		return err
	}

	err = p.Service.CreateCompany(context.TODO(), companyEvent.Company.ID)
	if err != nil {
		return err
	}

	return nil
}

func (p *KafkaProcessor) addEmployeeToCompany(msg *ckafka.Message) error {
	event := schema.NewCompanyEmployeeEvent()
	err := event.ParseJson(msg.Value)
	if err != nil {
		return err
	}

	err = p.Service.AddEmployeeToCompany(context.TODO(), event.CompanyID, event.EmployeeID)
	if err != nil {
		return err
	}

	return nil
}
