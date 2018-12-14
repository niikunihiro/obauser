package main

import (
	"context"
	pb "github.com/niixo/obauser/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type UserServer struct{}

func (us UserServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	return &pb.User{
		UserId:   uint32(req.UserId),
		UserName: "obauser",
	}, nil
}

func (us UserServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	return &pb.User{
		UserId:   1,
		UserName: req.UserName,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":18686")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserApiServer(s, &UserServer{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
