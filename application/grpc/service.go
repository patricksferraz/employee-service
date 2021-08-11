package grpc

import (
	"context"

	"github.com/c-4u/employee-service/application/grpc/pb"
	"github.com/c-4u/employee-service/domain/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GrpcService struct {
	pb.UnimplementedEmployeeServiceServer
	EmployeeService *service.Service
}

func NewGrpcService(service *service.Service) *GrpcService {
	return &GrpcService{
		EmployeeService: service,
	}
}

func (s *GrpcService) CreateEmployee(ctx context.Context, in *pb.CreateEmployeeRequest) (*pb.CreateEmployeeResponse, error) {
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

func (s *GrpcService) FindEmployee(ctx context.Context, in *pb.FindEmployeeRequest) (*pb.Employee, error) {
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

func (s *GrpcService) SetPassword(ctx context.Context, in *pb.SetPasswordRequest) (*pb.StatusResponse, error) {
	err := s.EmployeeService.SetPassword(ctx, in.EmployeeId, in.Password, in.Temporary)
	if err != nil {
		return &pb.StatusResponse{
			Code:  uint32(status.Code(err)),
			Error: err.Error(),
		}, err
	}
	return &pb.StatusResponse{
		Code:    uint32(codes.OK),
		Message: "password updated successfully",
	}, nil
}
