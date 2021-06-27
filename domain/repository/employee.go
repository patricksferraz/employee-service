package repository

import (
	"context"

	"dev.azure.com/c4ut/TimeClock/_git/employee-service/domain/entity"
)

type EmployeeRepositoryInterface interface {
	CreateEmployee(ctx context.Context, Employee *entity.Employee) error
	FindEmployee(ctx context.Context, id string) (*entity.Employee, error)
}
