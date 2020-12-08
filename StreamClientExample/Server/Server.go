package main

import (
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	generatedgrpc "stream-client-example/generated"
)

type server struct {
}

func (*server) LongGreet(stream generatedgrpc.GreetService_LongGreetServer) error {
	fmt.Printf("Long greet function was invoked with %v \n", stream)
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// Finished reading from the client stream.
			return stream.SendAndClose(&generatedgrpc.LongGreetResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Error reading client stream... %v", err)
		}

		firstName := req.GetGreeting().GetFirstName()
		result += "Hello " + firstName + "! "
	}
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
