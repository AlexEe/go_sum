// Package main implements a server for Greeter service.
package main

import (
	"context"
	"log"
	"net"
	"os"

	"goSum/pkg/proto"
	"goSum/pkg/sum"

	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type server struct{}

func (s *server) Sum(ctx context.Context, request *proto.SumRequest) (*proto.SumResult, error) {
	// Receive array of ints from request and add them up
	result, err := sum.Calculate(request.Numbers)
	if err != nil {
		os.Exit(1)
	}
	// Send back result in response
	return &proto.SumResult{Result: result}, nil
}

func main() {
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
