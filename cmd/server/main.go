package main

import (
	"log"
	"net"

	"github.com/mmatros/image-proc-api/api"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:12000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	api.RegisterImageProcApiServer(grpcServer, api.NewServer())
	grpcServer.Serve(lis)
}
