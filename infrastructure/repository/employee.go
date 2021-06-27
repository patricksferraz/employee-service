package repository

import (
	"context"

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
	createdAt := employee.CreatedAt.Unix()
	user := gocloak.User{
		ID:               &employee.ID,
		CreatedTimestamp: &createdAt,
		Username:         &employee.Username,
		FirstName:        &employee.FirstName,
		LastName:         &employee.LastName,
		Email:            &employee.Email,
		Enabled:          &employee.Enabled,
		EmailVerified:    &employee.EmailVerified,
	}
	user.Attributes = utils.StructToAttr(employee.Attributes)

	token, err := r.K.Client.LoginAdmin(ctx, r.K.Username, r.K.Password, r.K.Realm)
	if err != nil {
		return err
	}

	// TODO: check if session is ended
	defer r.K.Client.LogoutUserSession(ctx, token.AccessToken, r.K.Realm, token.SessionState)

	_, err = r.K.Client.CreateUser(ctx, token.AccessToken, r.K.Realm, user)
	if err != nil {
		return err
	}

	return nil
}
