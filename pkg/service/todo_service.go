package service

import (
    pb "github.com/Anant-1209/grpc-go/pkg/api/v1"
    "context"
)

type TodoService struct {
    pb.UnimplementedTodoServiceServer
}

func (s *TodoService) CreateTodo(ctx context.Context, req *pb.CreateTodoRequest) (*pb.CreateTodoResponse, error) {
    return &pb.CreateTodoResponse{Id: "123"}, nil
}
