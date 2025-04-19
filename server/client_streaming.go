package main

import (
	"io"
	"log"

	pb "github.com/coziamshreyansh/go-grpc-project/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var receivedMessages []string

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// Client has finished sending messages
			return stream.SendAndClose(&pb.MessageList{Messages: receivedMessages})
		}
		if err != nil {
			log.Printf("error receiving from stream: %v", err)
			return err
		}

		log.Printf("Received message: %v", req.GetMessage())
		receivedMessages = append(receivedMessages, req.GetMessage())
	}
}
