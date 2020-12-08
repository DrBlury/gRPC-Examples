package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	generatedgrpc "stream-server-example/generated"
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

	doServerStreaming(c)
}

func doServerStreaming(c generatedgrpc.GreetServiceClient) {
	fmt.Printf("Starting to do server streaming RPC...")

	req := &generatedgrpc.GreetManyTimesRequest{
		Greeting: &generatedgrpc.Greeting{
			FirstName: "Peter",
			LastName:  "Griffin",
		},
	}
	resultStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling server streaming GreetManyTimes RPC... %v", err)
	}

	for {
		msg, err := resultStream.Recv()
		if err == io.EOF {
			// Reached the end of the stream
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream...%v", err)
		}

		log.Printf("Response from GreetManyTimes: %v", msg.GetResult())
	}
}
