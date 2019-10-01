package sumCmd

import (
	"context"
	"fmt"
	"goSum/pkg/proto"
	"log"
	"time"

	"github.com/spf13/cobra"
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

// rootCmd represents the base command when called without any subcommands
var sumCmd = &cobra.Command{
	Use:     "sum",
	Short:   "Add numbers entered on the Command Line",
	Example: "sum -n 4,1,-2 -u localhost:8080",
	Run: func(cmd *cobra.Command, args []string) {
		// If address has been set via flag use input
		// Else use default address
		if address == "" {
			address = addressDefault
		}

		// Throw error if sum command is entered without specifying numbers to be added
		if len(numbers) < 1 {
			log.Fatalf("No numbers were entered. Example command: 'sum -n 1,3,4'")
		}

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
		client := proto.NewSumServiceClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		// New client makes request to server with array received through CLI
		// Server returns result of calculation
		result, err := client.Sum(ctx, &proto.SumRequest{Numbers: numbers32})
		if err != nil {
			log.Fatalf("Could not sum: %v", err)
		}

		// Print out result of server calculation
		fmt.Printf("The sum of ")
		for i, v := range numbers {
			if i == len(numbers)-2 {
				fmt.Print(v, " and ")
			} else if i == len(numbers)-1 {
				fmt.Print(v, " ")
			} else {
				fmt.Print(v, ", ")
			}
		}
		fmt.Printf("is %v.\n", result.GetResult())
	},
}

func init() {
	sumCmd.Flags().IntSliceVarP(&numbers, "numbers", "n", []int{}, "Numbers to be added up")
	sumCmd.Flags().StringVarP(&address, "url", "u", "", "")
	// sumCmd.AddSubCommands(rootCmd)
}

// AddSubCommands adds the sub-commands to the provided command
func AddSubCommands(cmd *cobra.Command) {
	cmd.AddCommand(sumCmd)
}
