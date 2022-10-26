package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Maverickme222222/emails/emailmgmt"
	"google.golang.org/grpc"
)

const (
	address = "localhost:9090"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect")
	}
	defer conn.Close()

	c := pb.NewEmailManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.CreateNewEmail(ctx, &pb.NewEmail{
		Name: "Email",
	})
	if err != nil {
		log.Fatalf("could not create email: %v", err)
	}

	log.Printf(`Email details
	Name: %s
	`, r.GetName())

}
