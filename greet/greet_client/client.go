package main

import (
	"context"
	"fmt"
	"go/gRPC/thomas/greet/greetpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("client is running")
	// create the connection
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	// close the connection when done
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Thomas",
			LastName:  "Greenhalgh",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC : %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}
