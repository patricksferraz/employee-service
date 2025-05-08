package repository

import (
	"context"

	"github.com/patricksferraz/employee-service/domain/entity"
	"github.com/patricksferraz/employee-service/domain/entity/filter"
)

type RepositoryInterface interface {
	CreateEmployee(ctx context.Context, Employee *entity.Employee) error
	FindEmployee(ctx context.Context, id string) (*entity.Employee, error)
	SearchEmployees(ctx context.Context, filter *filter.EmployeeFilter) (*string, []*entity.Employee, error)
	SaveEmployee(ctx context.Context, employee *entity.Employee) error

	PublishEvent(ctx context.Context, msg, topic, key string) error

	CreateUser(ctx context.Context, user *entity.User) error

	CreateCompany(ctx context.Context, company *entity.Company) error
	FindCompany(ctx context.Context, id string) (*entity.Company, error)
}
