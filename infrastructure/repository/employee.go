package repository

import (
	"context"
	"time"

	"github.com/Nerzal/gocloak/v8"
	"github.com/c-4u/employee-service/domain/entity"
	"github.com/c-4u/employee-service/infrastructure/external"
	"github.com/c-4u/employee-service/utils"
)

type KeycloakEmployeeRepository struct {
	K *external.Keycloak
}

func NewKeycloakEmployeeRepository(keycloak *external.Keycloak) *KeycloakEmployeeRepository {
	return &KeycloakEmployeeRepository{
		K: keycloak,
	}
}

func (r *KeycloakEmployeeRepository) CreateEmployee(ctx context.Context, employee *entity.Employee) (*string, error) {
	user := gocloak.User{
		Username:      &employee.Username,
		FirstName:     &employee.FirstName,
		LastName:      &employee.LastName,
		Email:         &employee.Email,
		Enabled:       &employee.Enabled,
		EmailVerified: &employee.EmailVerified,
	}
	user.Attributes = utils.StructToAttr(employee)

	token, err := r.K.Client.LoginAdmin(ctx, r.K.Username, r.K.Password, r.K.Realm)
	if err != nil {
		return nil, err
	}
	defer r.K.Client.LogoutUserSession(ctx, token.AccessToken, r.K.Realm, token.SessionState)

	employeeID, err := r.K.Client.CreateUser(ctx, token.AccessToken, r.K.Realm, user)
	if err != nil {
		return nil, err
	}

	return &employeeID, nil
}

func (r *KeycloakEmployeeRepository) FindEmployee(ctx context.Context, id string) (*entity.Employee, error) {
	token, err := r.K.Client.LoginAdmin(ctx, r.K.Username, r.K.Password, r.K.Realm)
	if err != nil {
		return nil, err
	}
	defer r.K.Client.LogoutUserSession(ctx, token.AccessToken, r.K.Realm, token.SessionState)

	e, err := r.K.Client.GetUserByID(ctx, token.AccessToken, r.K.Realm, id)
	if err != nil {
		return nil, err
	}

	var pis string
	if e.Attributes != nil {
		pis = (*e.Attributes)["pis"][0]
	}

	employee := &entity.Employee{
		Username:      *e.Username,
		FirstName:     *e.FirstName,
		LastName:      *e.LastName,
		Email:         *e.Email,
		Enabled:       *e.Enabled,
		EmailVerified: *e.EmailVerified,
		Pis:           pis,
	}
	employee.ID = *e.ID
	employee.CreatedAt = time.Unix(0, *e.CreatedTimestamp*int64(time.Millisecond))

	return employee, nil
}

func (r *KeycloakEmployeeRepository) SetPassword(ctx context.Context, employeeID string, pass *entity.PasswordInfo) error {
	token, err := r.K.Client.LoginAdmin(ctx, r.K.Username, r.K.Password, r.K.Realm)
	if err != nil {
		return err
	}
	defer r.K.Client.LogoutUserSession(ctx, token.AccessToken, r.K.Realm, token.SessionState)

	err = r.K.Client.SetPassword(ctx, token.AccessToken, employeeID, r.K.Realm, pass.Password, pass.Temporary)
	if err != nil {
		return err
	}

	return nil
}

func (r *KeycloakEmployeeRepository) SearchEmployees(ctx context.Context, filter *entity.Filter) ([]*entity.Employee, error) {
	token, err := r.K.Client.LoginAdmin(ctx, r.K.Username, r.K.Password, r.K.Realm)
	if err != nil {
		return nil, err
	}
	defer r.K.Client.LogoutUserSession(ctx, token.AccessToken, r.K.Realm, token.SessionState)

	first := filter.Page * filter.PageSize
	users, err := r.K.Client.GetUsers(
		ctx,
		token.AccessToken,
		r.K.Realm,
		gocloak.GetUsersParams{
			FirstName: &filter.FirstName,
			LastName:  &filter.LastName,
			First:     &first,
			Max:       &filter.PageSize,
		},
	)
	if err != nil {
		return nil, err
	}

	var employees []*entity.Employee
	for _, u := range users {
		var pis string
		if u.Attributes != nil {
			pis = (*u.Attributes)["pis"][0]
		}
		employee := &entity.Employee{
			Username:      *u.Username,
			FirstName:     *u.FirstName,
			LastName:      *u.LastName,
			Email:         *u.Email,
			Pis:           pis,
			Enabled:       *u.Enabled,
			EmailVerified: *u.EmailVerified,
		}
		employee.ID = *u.ID
		employee.CreatedAt = time.Unix(0, *u.CreatedTimestamp*int64(time.Millisecond))
		employees = append(employees, employee)
	}

	return employees, nil
}

func (r *KeycloakEmployeeRepository) UpdateEmployee(ctx context.Context, employee *entity.Employee) error {
	token, err := r.K.Client.LoginAdmin(ctx, r.K.Username, r.K.Password, r.K.Realm)
	if err != nil {
		return err
	}

	user := gocloak.User{
		ID:        &employee.ID,
		FirstName: &employee.FirstName,
		LastName:  &employee.LastName,
		Email:     &employee.Email,
	}
	user.Attributes = utils.StructToAttr(employee)

	err = r.K.Client.UpdateUser(ctx, token.AccessToken, r.K.Realm, user)
	if err != nil {
		return err
	}

	return nil
}
