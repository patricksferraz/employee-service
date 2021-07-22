package service

import (
	"context"

	"github.com/c-4u/employee-service/domain/entity"
	"github.com/c-4u/employee-service/domain/repository"
)

type EmployeeService struct {
	EmployeeRepository repository.EmployeeRepositoryInterface
}

func NewEmployeeService(employeeRepository repository.EmployeeRepositoryInterface) *EmployeeService {
	return &EmployeeService{
		EmployeeRepository: employeeRepository,
	}
}

func (e *EmployeeService) CreateEmployee(ctx context.Context, username, firstName, lastName, email, pis string, enabled, emailVerified bool) (*string, error) {
	employee, err := entity.NewEmployee("", username, firstName, lastName, email, pis, enabled, emailVerified)
	if err != nil {
		return nil, err
	}

	id, err := e.EmployeeRepository.CreateEmployee(ctx, employee)
	if err != nil {
		return nil, err
	}

	return id, nil
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

func (s *EmployeeService) SearchEmployees(ctx context.Context, firstName, lastName string, pageSize int, pageItems int) ([]*entity.Employee, error) {
	filter, err := entity.NewFilter(firstName, lastName, pageSize, pageItems)
	if err != nil {
		return nil, err
	}

	employees, err := s.EmployeeRepository.SearchEmployees(ctx, filter)
	if err != nil {
		return nil, err
	}

	return employees, nil
}

func (e *EmployeeService) UpdateEmployee(ctx context.Context, id, firstName, lastName, email string) error {
	_e, err := e.EmployeeRepository.FindEmployee(ctx, id)
	if err != nil {
		return err
	}

	employee, err := entity.NewEmployee(_e.ID, _e.Username, firstName, lastName, email, _e.Pis, _e.Enabled, _e.EmailVerified)
	if err != nil {
		return err
	}

	err = e.EmployeeRepository.UpdateEmployee(ctx, employee)
	if err != nil {
		return err
	}

	return nil
}
