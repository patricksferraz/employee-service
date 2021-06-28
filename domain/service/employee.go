package service

import (
	"context"

	"dev.azure.com/c4ut/TimeClock/_git/employee-service/domain/entity"
	"dev.azure.com/c4ut/TimeClock/_git/employee-service/domain/repository"
)

type EmployeeService struct {
	EmployeeRepository repository.EmployeeRepositoryInterface
}

func NewEmployeeService(employeeRepository repository.EmployeeRepositoryInterface) *EmployeeService {
	return &EmployeeService{
		EmployeeRepository: employeeRepository,
	}
}

func (e *EmployeeService) CreateEmployee(ctx context.Context, username, firstName, lastName, email, pis string, enabled, emailVerified bool) (*entity.Employee, error) {
	employee, err := entity.NewEmployee(username, firstName, lastName, email, pis, enabled, emailVerified)
	if err != nil {
		return nil, err
	}

	err = e.EmployeeRepository.CreateEmployee(ctx, employee)
	if err != nil {
		return nil, err
	}

	return employee, nil
}

func (e *EmployeeService) FindEmployee(ctx context.Context, id string) (*entity.Employee, error) {
	employee, err := e.EmployeeRepository.FindEmployee(ctx, id)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

func (e *EmployeeService) SetPassword(ctx context.Context, employeeID string, password string, temporary bool) error {
	pass, err := entity.NewPasswordInfo(password, temporary)
	if err != nil {
		return err
	}

	err = e.EmployeeRepository.SetPassword(ctx, employeeID, pass)
	return err
}
