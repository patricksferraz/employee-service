package repository

import (
	"context"
	"fmt"
	"time"

	"dev.azure.com/c4ut/TimeClock/_git/employee-service/domain/entity"
	"dev.azure.com/c4ut/TimeClock/_git/employee-service/infrastructure/external"
	"dev.azure.com/c4ut/TimeClock/_git/employee-service/utils"
	"github.com/Nerzal/gocloak/v8"
)

type KeycloakEmployeeRepository struct {
	K *external.Keycloak
}

func NewKeycloakEmployeeRepository(keycloak *external.Keycloak) *KeycloakEmployeeRepository {
	return &KeycloakEmployeeRepository{
		K: keycloak,
	}
}

func (r *KeycloakEmployeeRepository) CreateEmployee(ctx context.Context, employee *entity.Employee) error {
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
		return err
	}
	defer r.K.Client.LogoutUserSession(ctx, token.AccessToken, r.K.Realm, token.SessionState)

	employeeID, err := r.K.Client.CreateUser(ctx, token.AccessToken, r.K.Realm, user)
	if err != nil {
		return err
	}
	employee.ID = employeeID

	return nil
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

	employee := &entity.Employee{
		Username:      *e.Username,
		FirstName:     *e.FirstName,
		LastName:      *e.LastName,
		Email:         *e.Email,
		Enabled:       *e.Enabled,
		EmailVerified: *e.EmailVerified,
		Pis:           (*e.Attributes)["pis"][0],
	}
	employee.ID = *e.ID
	fmt.Println(*e.CreatedTimestamp)
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
