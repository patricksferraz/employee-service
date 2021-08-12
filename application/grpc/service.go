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
	Service *service.Service
}

func NewGrpcService(service *service.Service) *GrpcService {
	return &GrpcService{
		Service: service,
	}
}

func (s *GrpcService) CreateEmployee(ctx context.Context, in *pb.CreateEmployeeRequest) (*pb.CreateEmployeeResponse, error) {
	id, err := s.Service.CreateEmployee(
		ctx,
		in.Employee.FirstName,
		in.Employee.LastName,
		in.Employee.Email,
		in.Employee.Pis,
		in.Employee.Cpf,
	)
	if err != nil {
		return &pb.CreateEmployeeResponse{}, err
	}

	return &pb.CreateEmployeeResponse{Id: *id}, nil
}

func (s *GrpcService) FindEmployee(ctx context.Context, in *pb.FindEmployeeRequest) (*pb.FindEmployeeResponse, error) {
	employee, err := s.Service.FindEmployee(ctx, in.Id)
	if err != nil {
		return &pb.FindEmployeeResponse{}, err
	}

	return &pb.FindEmployeeResponse{
		Employee: &pb.Employee{
			Id:        employee.ID,
			FirstName: employee.FirstName,
			LastName:  employee.LastName,
			Email:     employee.Email,
			Pis:       employee.Pis,
			Cpf:       employee.Cpf,
			Enabled:   employee.Enabled,
			CreatedAt: timestamppb.New(employee.CreatedAt),
			UpdatedAt: timestamppb.New(employee.UpdatedAt),
		},
	}, nil
}

func (s *GrpcService) SearchEmployees(ctx context.Context, in *pb.SearchEmployeesRequest) (*pb.SearchEmployeesResponse, error) {
	nextPageToken, employees, err := s.Service.SearchEmployees(ctx, in.Filter.FirstName, in.Filter.LastName, int(in.Filter.PageSize), in.Filter.PageToken)
	if err != nil {
		return &pb.SearchEmployeesResponse{}, err
	}

	var result []*pb.Employee
	for _, employee := range employees {
		result = append(
			result,
			&pb.Employee{
				Id:        employee.ID,
				FirstName: employee.FirstName,
				LastName:  employee.LastName,
				Email:     employee.Email,
				Pis:       employee.Pis,
				Cpf:       employee.Cpf,
				Enabled:   employee.Enabled,
				CreatedAt: timestamppb.New(employee.CreatedAt),
				UpdatedAt: timestamppb.New(employee.UpdatedAt),
			},
		)
	}

	return &pb.SearchEmployeesResponse{NextPageToken: *nextPageToken, Employees: result}, nil
}

func (s *GrpcService) UpdateEmployee(ctx context.Context, in *pb.UpdateEmployeeRequest) (*pb.StatusResponse, error) {
	err := s.Service.UpdateEmployee(ctx, in.Id, in.FirstName, in.LastName, in.Email)
	if err != nil {
		return &pb.StatusResponse{
			Code:    uint32(status.Code(err)),
			Message: "not updated",
			Error:   err.Error(),
		}, err
	}

	return &pb.StatusResponse{
		Code:    uint32(codes.OK),
		Message: "successfully updated",
	}, nil
}
