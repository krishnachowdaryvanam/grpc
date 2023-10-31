package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "grpc_assignment/users"

	"google.golang.org/grpc"
)

type User struct {
	ID      int32
	Fname   string
	City    string
	Phone   int64
	Height  float32
	Married bool
}

var userDB = make(map[int32]User)

func init() {
	// Mock user data
	userDB[1] = User{ID: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true}
	userDB[2] = User{ID: 2, Fname: "Alice", City: "NYC", Phone: 9876543210, Height: 5.6, Married: false}
	userDB[3] = User{ID: 3, Fname: "Bob", City: "Chicago", Phone: 5551234567, Height: 6.1, Married: true}
}

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) GetUserById(ctx context.Context, req *pb.UserRequest) (*pb.User, error) {
	user, found := userDB[req.Id]
	if !found {
		return nil, fmt.Errorf("User not found")
	}
	return &pb.User{
		Id:      user.ID,
		Fname:   user.Fname,
		City:    user.City,
		Phone:   user.Phone,
		Height:  user.Height,
		Married: user.Married,
	}, nil
}

func (s *server) GetUsersByIds(req *pb.UserIdsRequest, stream pb.UserService_GetUsersByIdsServer) error {
	for _, userID := range req.Ids {
		user, found := userDB[userID]
		if !found {
			continue
		}
		err := stream.Send(&pb.User{
			Id:      user.ID,
			Fname:   user.Fname,
			City:    user.City,
			Phone:   user.Phone,
			Height:  user.Height,
			Married: user.Married,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	fmt.Println("gRPC server is running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
