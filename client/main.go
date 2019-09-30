package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"goSum/pkg/proto"

	"google.golang.org/grpc"
)

const (
	address = "localhost:8080"
)

func main() {
	var numbers []int32

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := proto.NewSumServiceClient(conn)

	// Contact the server and print out its response.
	// if len(os.Args) > 1 {
	numbers = []int32{1, 3, 3}
	// }
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := client.Sum(ctx, &proto.SumRequest{Numbers: numbers})
	if err != nil {
		log.Fatalf("Could not sum: %v", err)
	}
	fmt.Printf("The sum of ")
	for _, v := range numbers {
		fmt.Print(v, " ")
	}
	fmt.Printf("is %v.\n", result.GetResult())
}
