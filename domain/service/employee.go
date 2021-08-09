package service

import (
	"context"

	"github.com/c-4u/employee-service/domain/entity"
	"github.com/c-4u/employee-service/domain/repository"
	"github.com/c-4u/employee-service/infrastructure/external/topic"
)

type EmployeeService struct {
	EmployeeRepository repository.EmployeeRepositoryInterface
	EventRepository    repository.EventRepositoryInterface
}

func NewEmployeeService(employeeRepository repository.EmployeeRepositoryInterface, eventRepository repository.EventRepositoryInterface) *EmployeeService {
	return &EmployeeService{
		EmployeeRepository: employeeRepository,
		EventRepository:    eventRepository,
	}
}

func (s *EmployeeService) CreateEmployee(ctx context.Context, username, firstName, lastName, email, pis string, enabled, emailVerified bool) (*string, error) {
	employee, err := entity.NewEmployee("", username, firstName, lastName, email, pis, enabled, emailVerified)
	if err != nil {
		return nil, err
	}

	err = s.EmployeeRepository.CreateEmployee(ctx, employee)
	if err != nil {
		return nil, err
	}

	event, err := entity.NewEvent(employee)
	if err != nil {
		return nil, err
	}

	msg, err := event.ToJson()
	if err != nil {
		return nil, err
	}

	err = s.EventRepository.Publish(ctx, string(msg), topic.Employees, employee.ID)
	if err != nil {
		return nil, err
	}

	return &employee.ID, nil
}

func (s *EmployeeService) FindEmployee(ctx context.Context, id string) (*entity.Employee, error) {
	employee, err := s.EmployeeRepository.FindEmployee(ctx, id)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

func (s *EmployeeService) SetPassword(ctx context.Context, employeeID string, password string, temporary bool) error {
	pass, err := entity.NewPasswordInfo(password, temporary)
	if err != nil {
		return err
	}

	err = s.EmployeeRepository.SetPassword(ctx, employeeID, pass)
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

func (s *EmployeeService) UpdateEmployee(ctx context.Context, id, firstName, lastName, email string) error {
	e, err := s.EmployeeRepository.FindEmployee(ctx, id)
	if err != nil {
		return err
	}

	employee, err := entity.NewEmployee(e.ID, e.Username, firstName, lastName, email, e.Pis, e.Enabled, e.EmailVerified)
	if err != nil {
		return err
	}

	err = s.EmployeeRepository.UpdateEmployee(ctx, employee)
	if err != nil {
		return err
	}

	event, err := entity.NewEvent(employee)
	if err != nil {
		return err
	}

	msg, err := event.ToJson()
	if err != nil {
		return err
	}

	err = s.EventRepository.Publish(ctx, string(msg), topic.Employees, employee.ID)
	if err != nil {
		return err
	}

	return nil
}
