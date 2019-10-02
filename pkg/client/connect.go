package client

import (
	"context"
	"errors"
	"fmt"
	"goSum/pkg/proto"
	"log"
	"regexp"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"google.golang.org/grpc"
)

const (
	addressDefault = "localhost:8080"
)

var (
	numbers   []int
	numbers32 []int32
	address   string
)

// Connect opens a gRPC connection to the server and returns sum of numbers
func Connect(numbers []int, address string) {
	// If address has been set via flag use input
	// Else use default address
	if address == "" {
		address = addressDefault
	}

	// Throw error if sum command is entered without specifying numbers to be added
	if len(numbers) < 1 {
		numbers = promptNumbers()
		// log.Fatalf("No numbers were entered. Example command: 'sum -n 1,3,4'")
	}

	// Convert ints to int32
	for _, v := range numbers {
		n := int32(v)
		numbers32 = append(numbers32, n)
	}

	// Set up a connection to the server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	// Create a new client for the sum service
	client := proto.NewSumServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// New client makes request to server with array received through CLI
	// Server returns result of calculation
	result, err := client.Sum(ctx, &proto.SumRequest{Numbers: numbers32})
	if err != nil {
		log.Fatalf("Could not sum: %v", err)
	}

	printResult(numbers, result)
}

// printResult prints result of calculation on the CLI
func printResult(numbers []int, result *proto.SumResult) {
	whiteBold := color.New(color.FgHiWhite, color.Bold)

	switch len(numbers) {
	case 0:
		whiteBold.Println("\nNo numbers were added.")
	case 1:
		whiteBold.Printf("\nOnly one number was provided: %v \n", numbers[0])
	default:
		// Print out result of server calculation
		whiteBold.Printf("\nThe sum of ")
		for i, v := range numbers {
			if i == len(numbers)-2 {
				whiteBold.Print(v, " and ")
			} else if i == len(numbers)-1 {
				whiteBold.Print(v, " ")
			} else {
				whiteBold.Print(v, ", ")
			}
		}
		whiteBold.Printf("is %v.\n", result.GetResult())
	}
}

// promptNumbers asks user for input
func promptNumbers() []int {
	var next bool
	var numbers []int
	next = true

	for next == true {
		// Check that input is an int
		validate := func(input string) error {
			isInt := regexp.MustCompile("^[0-9]+$").MatchString
			switch isInt(input) {
			case true:
				return nil
			default:
				return errors.New("Invalid number")
			}
		}

		prompt := promptui.Prompt{
			Label:    "Enter numbers to be added ",
			Validate: validate,
		}

		result, err := prompt.Run()

		n, err := strconv.ParseInt(result, 36, 32)

		number := int(n)

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return nil
		}

		numbers = append(numbers, number)

		next = promptContinue()
	}

	return numbers
}

// promptContinue asks user is they want to keep adding numbers
func promptContinue() bool {
	prompt := promptui.Select{
		Label: "Add another number?",
		Items: []string{"Yes", "No"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	if result == "Yes" {
		return true
	} else {
		return false
	}

}
