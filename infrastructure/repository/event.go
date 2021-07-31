package repository

import (
	"context"

	"github.com/c-4u/employee-service/infrastructure/external"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaRepository struct {
	K *external.Kafka
}

func NewKafkaRepository(kafka *external.Kafka) *KafkaRepository {
	return &KafkaRepository{K: kafka}
}

func (r *KafkaRepository) Publish(ctx context.Context, msg, topic, key string) error {
	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		Value:          []byte(msg),
		Key:            []byte(key),
	}
	err := r.K.Producer.Produce(message, r.K.DeliveryChan)
	if err != nil {
		return err
	}
	return nil
}
