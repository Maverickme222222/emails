package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"

	pb "github.com/Maverickme222222/emails/emailmgmt"
	"google.golang.org/grpc"
)

const (
	port = ":9090"
)

type EmailManagementServer struct {
	pb.UnimplementedEmailManagementServer
}

func (s *EmailManagementServer) CreateNewEmail(ctx context.Context, req *pb.NewEmail) (*pb.NewEmailResponse, error) {
	log.Printf("Received: %v", req.Name)
	var email_id int32 = int32(rand.Intn(1000))
	email := fmt.Sprintf("Added new email id %v for name %v", email_id, req.GetName())
	return &pb.NewEmailResponse{
		Name: email,
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen at port %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterEmailManagementServer(s, &EmailManagementServer{})
	log.Printf("Server listening at %v", listen.Addr().String())
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
