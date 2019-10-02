package server

import (
	"context"
	"fmt"
	"goSum/pkg/calc/sum"
	"goSum/pkg/proto"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

const (
	portDefault = "8080"
)

var (
	port string
)

type server struct{}

// Sum receives the client request through gRPC and returns the result of the
// calculation
func (s *server) Sum(ctx context.Context, request *proto.SumRequest) (*proto.SumResult, error) {
	// Receive array of ints from request and add them up
	result, err := sum.Calculate(request.Numbers)
	if err != nil {
		os.Exit(1)
	}
	// Send back result in response
	return &proto.SumResult{Result: result}, nil
}

// Start starts a new server on default or provided port
func Start(port string) {
	if port == "" {
		port = portDefault
	}
	log.Printf("Starting new server on port %v.\n", port)

	port = fmt.Sprintf(":%v", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterSumServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
