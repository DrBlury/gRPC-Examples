package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	generatedgrpc "stream-server-example/generated"
	"time"
)

type server struct {
}

func (*server) GreetManyTimes(req *generatedgrpc.GreetManyTimesRequest, stream generatedgrpc.GreetService_GreetManyTimesServer) error {
	fmt.Printf("Greet many times function was invoked with %v", req)

	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hello! " + firstName + " number " + strconv.Itoa(i)
		res := &generatedgrpc.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
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
