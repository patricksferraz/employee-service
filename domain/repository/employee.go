package repository

import (
	"context"

	"dev.azure.com/c4ut/TimeClock/_git/employee-service/domain/entity"
)

type EmployeeRepositoryInterface interface {
	CreateEmployee(ctx context.Context, Employee *entity.Employee) (*string, error)
	FindEmployee(ctx context.Context, id string) (*entity.Employee, error)
	SetPassword(ctx context.Context, employeeID string, pass *entity.PasswordInfo) error
	SearchEmployees(ctx context.Context, filter *entity.Filter) ([]*entity.Employee, error)
	UpdateEmployee(ctx context.Context, employee *entity.Employee) error
}
