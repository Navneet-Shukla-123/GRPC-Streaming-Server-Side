package main

import (
	"grpc-streaming/file"
	"log"
	"net"

	"google.golang.org/grpc"
)

type streamingServer struct {
	file.UnimplementedMyStreamingServiceServer
}

// StreamData function will stream the total divisor of the number
func (s streamingServer) StreamData(req *file.RequestBody, stream file.MyStreamingService_StreamDataServer) error {

	reqBody := req.GetX()

	for i := 1; i <= int(reqBody); i++ {
		if int(reqBody)%i == 0 {

			resp := &file.ResponseBody{
				X: int32(i),
			}

			err := stream.Send(resp)
			if err != nil {
				log.Println("Error in sending the response")
				return err
			}

		}
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Error in establishing the tcp connection: %v", err)
	}

	grpcServer := grpc.NewServer()
	streaming := &streamingServer{}

	file.RegisterMyStreamingServiceServer(grpcServer, streaming)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Error in starting the grpc server: %v", err)
	}
}
