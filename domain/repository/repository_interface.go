package repository

import (
	"context"

	"github.com/c-4u/employee-service/domain/entity"
)

type RepositoryInterface interface {
	CreateEmployee(ctx context.Context, Employee *entity.Employee) error
	FindEmployee(ctx context.Context, id string) (*entity.Employee, error)
	SetPassword(ctx context.Context, employeeID string, pass *entity.PasswordInfo) error
	SearchEmployees(ctx context.Context, filter *entity.Filter) ([]*entity.Employee, error)
	UpdateEmployee(ctx context.Context, employee *entity.Employee) error

	PublishEmployeeEvent(ctx context.Context, msg, topic, key string) error
}
