package main

import (
	"context"
	pb "example.com/protobuf-example-go"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
)

const (
	port = ":50051"
)

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
}

func (s *UserManagementServer) CreateNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("The name of this user: %v", in.GetName())
	var user_id int32 = int32(rand.Intn(1000))
	return &pb.User{Name: in.GetName(), Age: in.GetAge(), Id: user_id}, nil
}

func main() {
	s := grpc.NewServer()
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}
	log.Printf("Currently listening at localhost%v", listen.Addr())
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve %v:", err)
	}
}
