package repository

import (
	"context"

	"github.com/c-4u/employee-service/domain/entity"
)

type RepositoryInterface interface {
	CreateEmployee(ctx context.Context, Employee *entity.Employee) error
	FindEmployee(ctx context.Context, id string) (*entity.Employee, error)
	SearchEmployees(ctx context.Context, filter *entity.Filter) (*string, []*entity.Employee, error)
	SaveEmployee(ctx context.Context, employee *entity.Employee) error

	PublishEvent(ctx context.Context, msg, topic, key string) error
}
