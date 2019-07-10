package main

import (
	"log"
	"net"

	pbf "./pbf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":8488"
)

// helloServer implements helloworld.HelloServer
type helloServer struct{}

// SayHello implements helloworld.HelloServer
func (s *helloServer) SayHello(ctx context.Context, in *pbf.HelloRequest) (*pbf.HelloResponse, error) {
	return &pbf.HelloResponse{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pbf.RegisterHelloServer(s, &helloServer{})

	// Register reflection service on gRPC server.
	// https://github.com/grpc/grpc-go/blob/master/Documentation/server-reflection-tutorial.md
	reflection.Register(s)

	log.Printf("Server started at port %s", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
