package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	generatedgrpc "stream-server-example/generated"
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

	doBidirectional(c)
}

func doBidirectional(c generatedgrpc.GreetServiceClient) {
	fmt.Println("Starting to do a Bi directional streaming RPC...")

	// Create stream by invoking the client and
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating the stream: %v", err)
		return
	}

	// Create a channel to block until everything is done
	waitchannel := make(chan struct{})

	// Send messages to client (own go routine)
	go sendRequests(stream)

	// Receive a lot of messages (own go routine)
	go receiveResponses(stream, waitchannel)

	// Block until everything is done
	<-waitchannel
}

func receiveResponses(stream generatedgrpc.GreetService_GreetEveryoneClient, waitchannel chan struct{}) {
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while receiving: %v", err)
			break
		}

		fmt.Printf("Received: " + res.GetResult())
	}
	close(waitchannel)
}

func sendRequests(stream generatedgrpc.GreetService_GreetEveryoneClient) {
	// Sends some messages
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
		stream.Send(&generatedgrpc.GreetEveryoneRequest{
			Greeting: &generatedgrpc.Greeting{
				FirstName: person.FirstName,
				LastName:  person.LastName,
			},
		})
		fmt.Println("\nSent over: " + person.FirstName)
		time.Sleep(1000 * time.Millisecond)
	}

	stream.CloseSend()
}
