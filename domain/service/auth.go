package service

import (
	"context"

	"github.com/patricksferraz/employee-service/application/grpc/pb"
	"github.com/patricksferraz/employee-service/domain/entity"
	"google.golang.org/grpc"
)

type AuthService struct {
	service pb.AuthKeycloakAclClient
}

func NewAuthService(cc *grpc.ClientConn) *AuthService {
	return &AuthService{
		service: pb.NewAuthKeycloakAclClient(cc),
	}
}

func (s *AuthService) Verify(ctx context.Context, accessToken string) (*entity.Claims, error) {
	req := &pb.FindClaimsByTokenRequest{
		AccessToken: accessToken,
	}

	_claims, err := s.service.FindClaimsByToken(ctx, req)
	if err != nil {
		return nil, err
	}

	claims, err := entity.NewClaims(_claims.EmployeeId, _claims.Roles)
	if err != nil {
		return nil, err
	}

	return claims, nil
}
