package grpc_server

import (
	"context"
	"greet_grpc_app/grpc_server/proto"
)

type server struct {
}

func (s server) Greet(ctx context.Context, request *proto.GreetRequest) (*proto.GreetResponse, error) {
	return &proto.GreetResponse{Message: "hello"}, nil
}

func New() proto.GreetServiceServer {
	return &server{}
}
