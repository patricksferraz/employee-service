package repository

import (
	"context"
	"time"

	"github.com/Nerzal/gocloak/v8"
	"github.com/c-4u/employee-service/domain/entity"
	"github.com/c-4u/employee-service/infrastructure/external"
	"github.com/c-4u/employee-service/utils"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type Repository struct {
	Keycloak *external.Keycloak
	Kafka    *external.Kafka
}

func NewRepository(keycloak *external.Keycloak, kafka *external.Kafka) *Repository {
	return &Repository{
		Keycloak: keycloak,
		Kafka:    kafka,
	}
}

func (r *Repository) CreateEmployee(ctx context.Context, employee *entity.Employee) error {
	user := gocloak.User{
		Username:      &employee.Username,
		FirstName:     &employee.FirstName,
		LastName:      &employee.LastName,
		Email:         &employee.Email,
		Enabled:       &employee.Enabled,
		EmailVerified: &employee.EmailVerified,
	}
	user.Attributes = utils.StructToAttr(employee)

	token, err := r.Keycloak.Client.LoginAdmin(ctx, r.Keycloak.Username, r.Keycloak.Password, r.Keycloak.Realm)
	if err != nil {
		return err
	}
	defer r.Keycloak.Client.LogoutUserSession(ctx, token.AccessToken, r.Keycloak.Realm, token.SessionState)

	employeeID, err := r.Keycloak.Client.CreateUser(ctx, token.AccessToken, r.Keycloak.Realm, user)
	if err != nil {
		return err
	}
	employee.ID = employeeID

	return nil
}

func (r *Repository) FindEmployee(ctx context.Context, id string) (*entity.Employee, error) {
	token, err := r.Keycloak.Client.LoginAdmin(ctx, r.Keycloak.Username, r.Keycloak.Password, r.Keycloak.Realm)
	if err != nil {
		return nil, err
	}
	defer r.Keycloak.Client.LogoutUserSession(ctx, token.AccessToken, r.Keycloak.Realm, token.SessionState)

	e, err := r.Keycloak.Client.GetUserByID(ctx, token.AccessToken, r.Keycloak.Realm, id)
	if err != nil {
		return nil, err
	}

	var pis string
	if e.Attributes != nil {
		pis = (*e.Attributes)["pis"][0]
	}

	employee := &entity.Employee{
		Username:      *e.Username,
		FirstName:     *e.FirstName,
		LastName:      *e.LastName,
		Email:         *e.Email,
		Enabled:       *e.Enabled,
		EmailVerified: *e.EmailVerified,
		Pis:           pis,
	}
	employee.ID = *e.ID
	employee.CreatedAt = time.Unix(0, *e.CreatedTimestamp*int64(time.Millisecond))

	return employee, nil
}

func (r *Repository) SetPassword(ctx context.Context, employeeID string, pass *entity.PasswordInfo) error {
	token, err := r.Keycloak.Client.LoginAdmin(ctx, r.Keycloak.Username, r.Keycloak.Password, r.Keycloak.Realm)
	if err != nil {
		return err
	}
	defer r.Keycloak.Client.LogoutUserSession(ctx, token.AccessToken, r.Keycloak.Realm, token.SessionState)

	err = r.Keycloak.Client.SetPassword(ctx, token.AccessToken, employeeID, r.Keycloak.Realm, pass.Password, pass.Temporary)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) SearchEmployees(ctx context.Context, filter *entity.Filter) ([]*entity.Employee, error) {
	token, err := r.Keycloak.Client.LoginAdmin(ctx, r.Keycloak.Username, r.Keycloak.Password, r.Keycloak.Realm)
	if err != nil {
		return nil, err
	}
	defer r.Keycloak.Client.LogoutUserSession(ctx, token.AccessToken, r.Keycloak.Realm, token.SessionState)

	first := filter.Page * filter.PageSize
	users, err := r.Keycloak.Client.GetUsers(
		ctx,
		token.AccessToken,
		r.Keycloak.Realm,
		gocloak.GetUsersParams{
			FirstName: &filter.FirstName,
			LastName:  &filter.LastName,
			First:     &first,
			Max:       &filter.PageSize,
		},
	)
	if err != nil {
		return nil, err
	}

	var employees []*entity.Employee
	for _, u := range users {
		var pis string
		if u.Attributes != nil {
			pis = (*u.Attributes)["pis"][0]
		}
		employee := &entity.Employee{
			Username:      *u.Username,
			FirstName:     *u.FirstName,
			LastName:      *u.LastName,
			Email:         *u.Email,
			Pis:           pis,
			Enabled:       *u.Enabled,
			EmailVerified: *u.EmailVerified,
		}
		employee.ID = *u.ID
		employee.CreatedAt = time.Unix(0, *u.CreatedTimestamp*int64(time.Millisecond))
		employees = append(employees, employee)
	}

	return employees, nil
}

func (r *Repository) UpdateEmployee(ctx context.Context, employee *entity.Employee) error {
	token, err := r.Keycloak.Client.LoginAdmin(ctx, r.Keycloak.Username, r.Keycloak.Password, r.Keycloak.Realm)
	if err != nil {
		return err
	}

	user := gocloak.User{
		ID:        &employee.ID,
		FirstName: &employee.FirstName,
		LastName:  &employee.LastName,
		Email:     &employee.Email,
	}
	user.Attributes = utils.StructToAttr(employee)

	err = r.Keycloak.Client.UpdateUser(ctx, token.AccessToken, r.Keycloak.Realm, user)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) PublishEmployeeEvent(ctx context.Context, msg, topic, key string) error {
	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		Value:          []byte(msg),
		Key:            []byte(key),
	}
	err := r.Kafka.Producer.Produce(message, r.Kafka.DeliveryChan)
	if err != nil {
		return err
	}
	return nil
}
