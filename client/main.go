package main

import (
	"log"

	pb "github.com/coziamshreyansh/go-grpc-project/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// both client and server are in same port but it doesnt matter at all client just need to know the port where server is listening
// Only the server defines the port it listens on; the client just connects to it.
const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Didnot connect %v", err)
	}

	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	callSayHello(client)
}
