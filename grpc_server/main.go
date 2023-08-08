package main

import (
	"fmt"
	"google.golang.org/grpc"
	"greet_grpc_app/grpc_server"
	"greet_grpc_app/grpc_server/proto"
	"log"
	"net"
)

var port int = 50051

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterGreetServiceServer(s, grpc_server.New())
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
