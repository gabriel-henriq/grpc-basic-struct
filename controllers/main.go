package controllers

import (
	"context"
	"grpc/app/pb"
)

type HelloServiceServer struct {
	Ctx context.Context
	pb.UnimplementedGreeterServer
}

func (s *HelloServiceServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello " + in.Name}, nil
}
