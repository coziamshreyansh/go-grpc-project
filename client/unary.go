package main

import (
	"context"
	"log"
	"time"

	pb "github.com/coziamshreyansh/go-grpc-project/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHelloClientStreaming(ctx)

	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}

	log.Printf("%s", res)

}
