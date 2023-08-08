package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "greet_grpc_app/proto"
	"log"
	"time"
)

var addr string = "localhost:50051"

func main() {
	// establish a connetion to the target server with - communication is not encrypted
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Unable to establish a connection with server: %v", err)
	}

	// defer execute before the surrouding function (main) returns. In reverse order (LIFO)
	defer conn.Close()

	// creates a new client
	c := pb.NewGreetServiceClient(conn);

	// what is context ?
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// calling method (service) as if it's local
	res, err := c.Greet(ctx, &pb.GreetRequest{})

	if err != nil {
		log.Fatalf("unable to call greet %v", err)
	}

	log.Printf("%s", res.GetMessage())
}
