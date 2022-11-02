package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"

	pb "github.com/Maverickme222222/emails/emailmgmt"
	"github.com/Maverickme222222/emails/health"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
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

	healthService := health.NewHealthChecker()
	grpc_health_v1.RegisterHealthServer(s, healthService)

	log.Printf("Emails Server listening at %v", listen.Addr().String())
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
