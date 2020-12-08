package main

import (
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	generatedgrpc "stream-server-example/generated"
)

type server struct {
}

func (*server) GreetEveryone(stream generatedgrpc.GreetService_GreetEveryoneServer) error {
	fmt.Printf("Greet everyone function was invoked\n")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}

		firstName := req.GetGreeting().GetFirstName()
		result := "Hello " + firstName
		err = stream.Send(&generatedgrpc.GreetEveryoneResponse{
			Result: result,
		})

		if err != nil {
			log.Fatalf("Error while sending data to client: %v", err)
			return err
		}
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
