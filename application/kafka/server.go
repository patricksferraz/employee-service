package kafka

import (
	"fmt"

	"github.com/patricksferraz/employee-service/domain/service"
	"github.com/patricksferraz/employee-service/infrastructure/db"
	"github.com/patricksferraz/employee-service/infrastructure/external"
	"github.com/patricksferraz/employee-service/infrastructure/repository"
)

func StartKafkaServer(database *db.Postgres, kafkaProducer *external.KafkaProducer, kafkaConsumer *external.KafkaConsumer) {
	repository := repository.NewRepository(database, kafkaProducer)
	service := service.NewService(repository)

	fmt.Println("kafka consumer has been started")
	processor := NewKafkaProcessor(service, kafkaConsumer)
	processor.Consume()
}
