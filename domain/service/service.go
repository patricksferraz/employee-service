package service

import (
	"context"

	"github.com/c-4u/employee-service/domain/entity"
	"github.com/c-4u/employee-service/domain/repository"
	"github.com/c-4u/employee-service/infrastructure/external/topic"
)

type Service struct {
	Repository repository.RepositoryInterface
}

func NewService(repository repository.RepositoryInterface) *Service {
	return &Service{
		Repository: repository,
	}
}

func (s *Service) CreateEmployee(ctx context.Context, username, firstName, lastName, email, pis string, enabled, emailVerified bool) (*string, error) {
	employee, err := entity.NewEmployee("", username, firstName, lastName, email, pis, enabled, emailVerified)
	if err != nil {
		return nil, err
	}

	err = s.Repository.CreateEmployee(ctx, employee)
	if err != nil {
		return nil, err
	}

	event, err := entity.NewEmployeeEvent(employee)
	if err != nil {
		return nil, err
	}

	msg, err := event.ToJson()
	if err != nil {
		return nil, err
	}

	err = s.Repository.PublishEmployeeEvent(ctx, string(msg), topic.NEW_EMPLOYEE, employee.ID)
	if err != nil {
		return nil, err
	}

	return &employee.ID, nil
}

func (s *Service) FindEmployee(ctx context.Context, id string) (*entity.Employee, error) {
	employee, err := s.Repository.FindEmployee(ctx, id)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

func (s *Service) SetPassword(ctx context.Context, employeeID string, password string, temporary bool) error {
	pass, err := entity.NewPasswordInfo(password, temporary)
	if err != nil {
		return err
	}

	err = s.Repository.SetPassword(ctx, employeeID, pass)
	return err
}

func (s *Service) SearchEmployees(ctx context.Context, firstName, lastName string, pageSize int, pageItems int) ([]*entity.Employee, error) {
	filter, err := entity.NewFilter(firstName, lastName, pageSize, pageItems)
	if err != nil {
		return nil, err
	}

	employees, err := s.Repository.SearchEmployees(ctx, filter)
	if err != nil {
		return nil, err
	}

	return employees, nil
}

func (s *Service) UpdateEmployee(ctx context.Context, id, firstName, lastName, email string) error {
	e, err := s.Repository.FindEmployee(ctx, id)
	if err != nil {
		return err
	}

	employee, err := entity.NewEmployee(e.ID, e.Username, firstName, lastName, email, e.Pis, e.Enabled, e.EmailVerified)
	if err != nil {
		return err
	}

	err = s.Repository.UpdateEmployee(ctx, employee)
	if err != nil {
		return err
	}

	event, err := entity.NewEmployeeEvent(employee)
	if err != nil {
		return err
	}

	msg, err := event.ToJson()
	if err != nil {
		return err
	}

	err = s.Repository.PublishEmployeeEvent(ctx, string(msg), topic.UPDATE_EMPLOYEE, employee.ID)
	if err != nil {
		return err
	}

	return nil
}
