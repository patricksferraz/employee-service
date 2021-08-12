package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/c-4u/employee-service/application/grpc/pb"
	_service "github.com/c-4u/employee-service/domain/service"
	"github.com/c-4u/employee-service/infrastructure/db"
	"github.com/c-4u/employee-service/infrastructure/external"
	"github.com/c-4u/employee-service/infrastructure/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGrpcServer(database *db.Postgres, authConn *grpc.ClientConn, kafka *external.Kafka, port int) {

	authService := _service.NewAuthService(authConn)
	interceptor := NewAuthInterceptor(authService)
	repository := repository.NewRepository(database, kafka)
	service := _service.NewService(repository)
	grpcService := NewGrpcService(service)

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	)
	reflection.Register(grpcServer)
	pb.RegisterEmployeeServiceServer(grpcServer, grpcService)

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
