package service

import (
	"context"

	"dev.azure.com/c4ut/TimeClock/_git/employee-service/application/grpc/pb"
	"dev.azure.com/c4ut/TimeClock/_git/employee-service/domain/entity"
)

type AuthService struct {
	Service pb.AuthServiceClient
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

func NewAuthService(service pb.AuthServiceClient) *AuthService {
	return &AuthService{
		Service: service,
	}
}
