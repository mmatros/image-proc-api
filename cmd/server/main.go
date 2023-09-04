package main

import (
	"log"
	"net"

	api "github.com/mmatros/image-proc-api/pkg/api/imageproc_v1"

	"github.com/mmatros/image-proc-api/internal/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:12000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	api.RegisterImageProcApiServer(grpcServer, server.NewServer())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("error on serve grpc %v", err)
	}
}
