package service

import (
	"context"

	"github.com/c-4u/employee-service/application/grpc/pb"
	"github.com/c-4u/employee-service/domain/entity"
)

type AuthService struct {
	Service pb.AuthKeycloakAclClient
}

func NewAuthService(service pb.AuthKeycloakAclClient) *AuthService {
	return &AuthService{
		Service: service,
	}
}

func (a *AuthService) Verify(ctx context.Context, accessToken string) (*entity.Claims, error) {
	req := &pb.FindClaimsByTokenRequest{
		AccessToken: accessToken,
	}

	_claims, err := a.Service.FindClaimsByToken(ctx, req)
	if err != nil {
		return nil, err
	}

	claims, err := entity.NewClaims(_claims.EmployeeId, _claims.Roles)
	if err != nil {
		return nil, err
	}

	return claims, nil
}
