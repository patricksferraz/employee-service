package grpc

import (
	"context"

	"dev.azure.com/c4ut/TimeClock/_git/employee-service/application/grpc/pb"
	"dev.azure.com/c4ut/TimeClock/_git/employee-service/domain/service"
)

type EmployeeGrpcService struct {
	pb.UnimplementedEmployeeServiceServer
	EmployeeService *service.EmployeeService
}

func NewEmployeeGrpcService(service *service.EmployeeService) *EmployeeGrpcService {
	return &EmployeeGrpcService{
		EmployeeService: service,
	}
}

func (s *EmployeeGrpcService) CreateEmployee(ctx context.Context, in *pb.CreateEmployeeRequest) (*pb.CreateEmployeeResponse, error) {
	id, err := s.EmployeeService.CreateEmployee(ctx, in.Username, in.FirstName, in.LastName, in.Email, in.Pis, in.Enabled, in.EmailVerified)
	if err != nil {
		return &pb.CreateEmployeeResponse{}, err
	}

	return &pb.CreateEmployeeResponse{Id: *id}, nil
}
