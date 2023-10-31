package main

import (
	"context"
	"fmt"
	"log"

	pb "grpc_assignment/users"

	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a client for the UserService
	client := pb.NewUserServiceClient(conn)

	// Test GetUserById
	user, err := client.GetUserById(context.Background(), &pb.UserRequest{Id: 1})
	if err != nil {
		log.Fatalf("Failed to call GetUserById: %v", err)
	}
	fmt.Printf("User by ID: %v\n", user)

	// Test GetUsersByIds
	ids := []int32{2, 3}
	stream, err := client.GetUsersByIds(context.Background(), &pb.UserIdsRequest{Ids: ids})
	if err != nil {
		log.Fatalf("Failed to call GetUsersByIds: %v", err)
	}
	for {
		user, err := stream.Recv()
		if err != nil {
			break
		}
		fmt.Printf("User by IDS from stream: %v\n", user)
	}
}
