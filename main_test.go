package main

import (
	"context"
	"testing"

	pb "grpc_assignment/users"

	"google.golang.org/grpc"
)

func TestGetUserById(t *testing.T) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to connect to the server: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	// Test case 1: Fetch a user by ID
	userID := int32(1)
	req := &pb.UserRequest{Id: userID}
	user, err := client.GetUserById(context.Background(), req)
	if err != nil {
		t.Fatalf("Failed to call GetUserById: %v", err)
	}

	if user == nil {
		t.Fatalf("User by ID is nil")
	}

	if user.Id != userID {
		t.Errorf("Returned user has incorrect ID. Expected: %d, Actual: %d", userID, user.Id)
	}
}

func TestGetUsersByIds(t *testing.T) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to connect to the server: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	// Test case 2: Fetch a list of users by their IDs
	ids := []int32{1, 2, 3}
	req := &pb.UserIdsRequest{Ids: ids}
	stream, err := client.GetUsersByIds(context.Background(), req)
	if err != nil {
		t.Fatalf("Failed to call GetUsersByIds: %v", err)
	}

	for i := 0; i < len(ids); i++ {
		user, err := stream.Recv()
		if err != nil {
			t.Fatalf("Failed to receive user from the stream: %v", err)
		}

		if user.Id != ids[i] {
			t.Errorf("Returned user has incorrect ID. Expected: %d, Actual: %d", ids[i], user.Id)
		}
	}
}
