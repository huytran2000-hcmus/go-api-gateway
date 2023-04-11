package client

import (
	"context"
	"os"

	pb "go-api-gateway/proto/auth"
	Grpc "go-api-gateway/util"
)

// Login client handler
func Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	conn := Grpc.Dial(os.Getenv("services_grpc"))
	defer conn.Close()

	client := pb.NewAuthClient(conn)

	return client.Login(ctx, req)
}
