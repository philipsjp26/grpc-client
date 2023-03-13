package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	pb "grpc_client/common"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to server: %v", err)
	}
	defer conn.Close()

	// Create a Greeter client from the connection.
	client := pb.NewGreeterClient(conn)

	// Call the SayHello method of the Greeter client.
	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "phil"})
	if err != nil {
		log.Fatalf("error calling SayHello: %v", err)
	}

	// Print the response from the server.
	fmt.Println(resp.GetMessage())
}
