package grpc

import (
	"fmt"
	"log"
	"net"

	"dev.azure.com/c4ut/TimeClock/_git/employee-service/application/grpc/pb"
	"dev.azure.com/c4ut/TimeClock/_git/employee-service/domain/service"
	"dev.azure.com/c4ut/TimeClock/_git/employee-service/infrastructure/external"
	"dev.azure.com/c4ut/TimeClock/_git/employee-service/infrastructure/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGrpcServer(keycloak *external.Keycloak, port int) {

	employeeRepository := repository.NewKeycloakEmployeeRepository(keycloak)
	employeeService := service.NewEmployeeService(employeeRepository)
	employeeGrpcService := NewEmployeeGrpcService(employeeService)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterEmployeeServiceServer(grpcServer, employeeGrpcService)

	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start grpc server", err)
	}

	log.Printf("gRPC server has been started on port %d", port)

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start grpc server", err)
	}
}
