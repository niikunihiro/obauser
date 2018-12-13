package main

import (
	"context"
	"github.com/labstack/gommon/log"
	pb "github.com/niixo/obauser/rpc"
	"google.golang.org/grpc"
	"time"
)

const (
	address = "localhost:18686"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserApiClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetUser(ctx, &pb.GetUserRequest{})
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}
	log.Printf("UserId: %d", r.UserId)
	log.Printf("UserName: %s", r.UserName)

	r, err = c.CreateUser(ctx, &pb.CreateUserRequest{UserName: "obauser"})
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}
	log.Printf("UserId: %d", r.UserId)
	log.Printf("UserName: %s", r.UserName)
}
