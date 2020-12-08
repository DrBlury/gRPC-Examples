package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	generatedgrpc "stream-client-example/generated"
	"time"
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

	//doUnary(c)

	//doServerStreaming(c)

	doClientStreaming(c)
}

func doClientStreaming(c generatedgrpc.GreetServiceClient) {
	fmt.Printf("Starting to do client streaming RPC...")

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet... %v", err)
	}

	type Person struct {
		FirstName string
		LastName  string
	}

	personArr := []Person{
		Person{FirstName: "John", LastName: "Snow"},
		Person{FirstName: "Peter", LastName: "Griffin"},
		Person{FirstName: "Johannes", LastName: "Heinz"},
		Person{FirstName: "Olaf", LastName: "Palmtree"},
		Person{FirstName: "Brigitte", LastName: "Superstar"},
		Person{FirstName: "Ygritte", LastName: "Schmidt"},
		Person{FirstName: "Nautilus", LastName: "Greiphon"},
	}

	for i := 0; i < len(personArr); i++ {
		person := personArr[i]
		stream.Send(&generatedgrpc.LongGreetRequest{
			Greeting: &generatedgrpc.Greeting{
				FirstName: person.FirstName,
				LastName:  person.LastName,
			},
		})
		fmt.Println("\nSent over: " + person.FirstName)
		time.Sleep(1000 * time.Millisecond)
	}
	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v", err)
	}

	fmt.Printf("LongGreet response: %v", response)
}
