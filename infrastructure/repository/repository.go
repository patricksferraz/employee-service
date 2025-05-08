package repository

import (
	"context"
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/patricksferraz/employee-service/domain/entity"
	"github.com/patricksferraz/employee-service/domain/entity/filter"
	"github.com/patricksferraz/employee-service/infrastructure/db"
	"github.com/patricksferraz/employee-service/infrastructure/external"
)

type Repository struct {
	P *db.Postgres
	K *external.KafkaProducer
}

func NewRepository(postgres *db.Postgres, kafkaProducer *external.KafkaProducer) *Repository {
	return &Repository{
		P: postgres,
		K: kafkaProducer,
	}
}

func (r *Repository) CreateEmployee(ctx context.Context, employee *entity.Employee) error {
	err := r.P.Db.Create(employee).Error
	return err
}

func (r *Repository) FindEmployee(ctx context.Context, id string) (*entity.Employee, error) {
	var employee entity.Employee
	r.P.Db.Preload("User").Preload("Companies").First(&employee, "id = ?", id)

	if employee.ID == "" {
		return nil, fmt.Errorf("no employee found")
	}

	return &employee, nil
}

func (r *Repository) SearchEmployees(ctx context.Context, filter *filter.EmployeeFilter) (*string, []*entity.Employee, error) {
	var employees []*entity.Employee

	q := r.P.Db.Order("token desc").Limit(filter.PageSize)

	if filter.FirstName != "" {
		q = q.Where("first_name = ?", filter.FirstName)
	}
	if filter.LastName != "" {
		q = q.Where("last_name = ?", filter.LastName)
	}
	if filter.PageToken != "" {
		q = q.Where("token < ?", filter.PageToken)
	}

	err := q.Preload("User").Preload("Companies").Find(&employees).Error
	if err != nil {
		return nil, nil, err
	}

	length := len(employees)
	var nextPageToken string
	if length == filter.PageSize {
		nextPageToken = *employees[length-1].Token
	}

	return &nextPageToken, employees, nil
}

func (r *Repository) SaveEmployee(ctx context.Context, employee *entity.Employee) error {
	err := r.P.Db.Save(employee).Error
	return err
}

func (r *Repository) PublishEvent(ctx context.Context, msg, topic, key string) error {
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

func (r *Repository) CreateUser(ctx context.Context, user *entity.User) error {
	err := r.P.Db.Create(user).Error
	return err
}

func (r *Repository) CreateCompany(ctx context.Context, company *entity.Company) error {
	err := r.P.Db.Create(company).Error
	return err
}

func (r *Repository) FindCompany(ctx context.Context, id string) (*entity.Company, error) {
	var company entity.Company
	r.P.Db.First(&company, "id = ?", id)

	if company.ID == "" {
		return nil, fmt.Errorf("no company found")
	}

	return &company, nil
}
