package main

import (
	"context"
	"log"

	"github.com/go-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()

	// use `WithInsecure` to make calls in plaintext (without TLS).
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	c := proto.NewHelloServiceClient(conn)

	r, err := c.SayHello(ctx, &proto.SayHelloRequest{Name: "Phelype"})
	if err != nil {
		log.Fatalf("Failed to greet: %v", err)
	}

	log.Printf("Successful response received: %s", r.GetMessage())
}
