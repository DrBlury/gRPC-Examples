package main

import (
	"context"
	"fmt"
	"log"
	generatedgrpc "protobuf-example/generated/gRPC"

	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Hello i'm the client.")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to server: %v", err)
	}

	defer cc.Close()

	c := generatedgrpc.NewGreetServiceClient(cc)
	fmt.Printf("Created client: %f", c)

	doUnary(c)
}

func doUnary(c generatedgrpc.GreetServiceClient) {
	fmt.Printf("Starting to do unary RPC...")
	req := &generatedgrpc.GreetRequest{
		Greeting: &generatedgrpc.Greeting{
			FirstName: "Peter",
			LastName:  "Griffin",
		},
	}

	greetResponse, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not call greeting: %v", err)
	}
	log.Printf("\n\nResponse from server: %v", greetResponse.Result)
}
