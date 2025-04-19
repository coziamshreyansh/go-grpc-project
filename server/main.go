package main

import (
	"log"
	"net"

	pb "github.com/coziamshreyansh/go-grpc-project/proto"
	"google.golang.org/grpc"
)

// define a port

const (
	port = ":8080"
)

type helloServer struct {
	pb.GreetServiceServer
	// is calling the generated function to register your helloServer as the gRPC server handler for the GreetService.
}

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}

	// create a grpc server
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("server start at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start grpc server %v", err)
	}
}
