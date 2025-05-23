package service

import (
	"context"

	"github.com/patricksferraz/employee-service/domain/entity"
	"github.com/patricksferraz/employee-service/domain/entity/event"
	"github.com/patricksferraz/employee-service/domain/entity/filter"
	"github.com/patricksferraz/employee-service/domain/repository"
	"github.com/patricksferraz/employee-service/infrastructure/external/topic"
)

type Service struct {
	Repository repository.RepositoryInterface
}

func NewService(repository repository.RepositoryInterface) *Service {
	return &Service{
		Repository: repository,
	}
}

func (s *Service) CreateEmployee(ctx context.Context, firstName, lastName, email, pis, cpf string) (*string, error) {
	employee, err := entity.NewEmployee(firstName, lastName, email, pis, cpf)
	if err != nil {
		return nil, err
	}

	err = s.Repository.CreateEmployee(ctx, employee)
	if err != nil {
		return nil, err
	}

	event, err := event.NewEmployeeEvent(employee)
	if err != nil {
		return nil, err
	}

	msg, err := event.ToJson()
	if err != nil {
		return nil, err
	}

	err = s.Repository.PublishEvent(ctx, string(msg), topic.NEW_EMPLOYEE, employee.ID)
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

func (s *Service) SearchEmployees(ctx context.Context, firstName, lastName string, pageSize int, pageToken string) (*string, []*entity.Employee, error) {
	filter, err := filter.NewEmployeeFilter(firstName, lastName, pageSize, pageToken)
	if err != nil {
		return nil, nil, err
	}

	nextPageToken, employees, err := s.Repository.SearchEmployees(ctx, filter)
	if err != nil {
		return nil, nil, err
	}

	return nextPageToken, employees, nil
}

func (s *Service) UpdateEmployee(ctx context.Context, id, firstName, lastName, email string) error {
	employee, err := s.Repository.FindEmployee(ctx, id)
	if err != nil {
		return err
	}

	if err = employee.SetFirstName(firstName); err != nil {
		return err
	}
	if err = employee.SetLastName(lastName); err != nil {
		return err
	}
	if err = employee.SetEmail(email); err != nil {
		return err
	}
	if err = s.Repository.SaveEmployee(ctx, employee); err != nil {
		return err
	}

	event, err := event.NewEmployeeEvent(employee)
	if err != nil {
		return err
	}

	msg, err := event.ToJson()
	if err != nil {
		return err
	}

	if err = s.Repository.PublishEvent(ctx, string(msg), topic.UPDATE_EMPLOYEE, employee.ID); err != nil {
		return err
	}

	return nil
}

func (s *Service) CreateUser(ctx context.Context, id, username, employeeID string) error {
	employee, err := s.Repository.FindEmployee(ctx, employeeID)
	if err != nil {
		return err
	}

	user, err := entity.NewUser(id, username, employee)
	if err != nil {
		return err
	}

	err = s.Repository.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) CreateCompany(ctx context.Context, id string) error {
	company, err := entity.NewCompany(id)
	if err != nil {
		return err
	}

	err = s.Repository.CreateCompany(ctx, company)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) AddEmployeeToCompany(ctx context.Context, companyID, employeeID string) error {
	company, err := s.Repository.FindCompany(ctx, companyID)
	if err != nil {
		return err
	}

	employee, err := s.Repository.FindEmployee(ctx, employeeID)
	if err != nil {
		return err
	}

	err = employee.AddCompany(company)
	if err != nil {
		return err
	}

	err = s.Repository.SaveEmployee(ctx, employee)
	if err != nil {
		return err
	}

	return nil
}
