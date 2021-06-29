package grpc

import (
	"context"

	"dev.azure.com/c4ut/TimeClock/_git/employee-service/application/grpc/pb"
	"dev.azure.com/c4ut/TimeClock/_git/employee-service/domain/service"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	id, err := s.EmployeeService.CreateEmployee(
		ctx,
		in.Employee.Username,
		in.Employee.FirstName,
		in.Employee.LastName,
		in.Employee.Email,
		in.Employee.Pis,
		in.Employee.Enabled,
		in.Employee.EmailVerified,
	)
	if err != nil {
		return &pb.CreateEmployeeResponse{}, err
	}

	return &pb.CreateEmployeeResponse{Id: *id}, nil
}

func (s *EmployeeGrpcService) FindEmployee(ctx context.Context, in *pb.FindEmployeeRequest) (*pb.Employee, error) {
	employee, err := s.EmployeeService.FindEmployee(ctx, in.EmployeeId)
	if err != nil {
		return &pb.Employee{}, err
	}

	return &pb.Employee{
		Id:            employee.ID,
		Username:      employee.Username,
		FirstName:     employee.FirstName,
		LastName:      employee.LastName,
		Email:         employee.Email,
		Pis:           employee.Pis,
		Enabled:       employee.Enabled,
		EmailVerified: employee.EmailVerified,
		CreatedAt:     timestamppb.New(employee.CreatedAt),
	}, nil
}
