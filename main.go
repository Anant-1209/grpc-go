package main

import (
	pb "github.com/Anant-1209/grpc-go/pkg/api/v1" // Update to match your module name
	"github.com/Anant-1209/grpc-go/pkg/service"  // Update to match your module name
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// Start the gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTodoServiceServer(s, &service.TodoService{})
	log.Println("Starting gRPC server on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
