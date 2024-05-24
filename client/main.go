package main

import (
	"context"
	"grpc-streaming/file"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Println("Error in connecting to grpc server ", err)
		return
	}

	c := file.NewMyStreamingServiceClient(conn)

	stream, err := c.StreamData(context.Background(), &file.RequestBody{
		X: 100000,
	})
	if err != nil {
		log.Println("Error in receiving the data from server ", err)
		return
	}
	for {
		resp, err := stream.Recv()
		if err != nil {
			// If end of stream
			if err == io.EOF {
				break
			}
			log.Fatalf("Error receiving response: %v", err)
		}

		log.Println("Received response is ", resp)
	}
}
