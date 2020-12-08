package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	generatedgrpc "protobuf-example/generated/gRPC"
)

type server struct {
}

func (*server) Greet(ctx context.Context, req *generatedgrpc.GreetRequest) (*generatedgrpc.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v", req)

	firstName := req.GetGreeting().GetFirstName()

	result := "Hello ufhidsfahuso" + firstName

	response := &generatedgrpc.GreetResponse{
		Result: result,
	}

	return response, nil
}

func main() {
	fmt.Println("Hello i'm the server.")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	generatedgrpc.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
