package external

import (
	"google.golang.org/grpc"
)

func GrpcClient(addr string) (*grpc.ClientConn, error) {

	conn, err := grpc.Dial(
		addr,
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
