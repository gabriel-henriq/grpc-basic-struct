package main

import (
	"fmt"
	"grpc/app/pb"
	"grpc/controllers"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", "50051"))
	if err != nil {
		log.Fatalf("Could not connect on port :%s: %s", "50051", err.Error())
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	service := controllers.HelloServiceServer{}

	pb.RegisterGreeterServer(grpcServer, &service)

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	fmt.Printf("Server started on port :%s\n", "50051")

	c := make(chan os.Signal)
	fmt.Printf("Waiting for SIGINT...\n")
	signal.Notify(c, os.Interrupt)
	fmt.Printf("Got SIGINT\n")
	<-c

	fmt.Printf("\nStopping the server...\n")
}
