package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/go-grpc/internal/hello"
	"github.com/go-grpc/proto"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer lis.Close()

	grpcServer := grpc.NewServer()

	// register our RPC implementation on the gRPC server.
	proto.RegisterHelloServiceServer(grpcServer, &hello.Server{})

	log.Println("Server is running on port 50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
