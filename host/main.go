package main

import (
	"fmt"
	"github.com/trendyol/smurfs/host/internal/service"
	"github.com/trendyol/smurfs/host/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:8080"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	reflection.Register(server)
	protos.RegisterLogServiceServer(server, service.NewLogService())
	if err = server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
