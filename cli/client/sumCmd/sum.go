package sumCmd

import (
	"fmt"
	"goSum/pkg/client"
	"io/ioutil"
	"log"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	numbers []int
	address string
)

// sumCmd defines the 'sum' subcommand
var sumCmd = &cobra.Command{
	Use:     "sum",
	Short:   "Add numbers entered on the Command Line",
	Example: "sum -n 4,1,-2 -u localhost:8080",
	Run: func(cmd *cobra.Command, args []string) {
		// Print out logo
		whiteBold := color.New(color.FgHiWhite, color.Bold)
		logo, err := ioutil.ReadFile("logo.txt")
		if err != nil {
			log.Fatalln("Error opening 'logo.txt':", err)
		}
		whiteBold.Println(string(logo))
		fmt.Print("Perform mathematical operations from the command line!\n\n")

		// Start gRPC connection to server
		client.Connect(numbers, address)
	},
}

// Add flags to allow input from the CLI
func init() {
	sumCmd.Flags().IntSliceVarP(&numbers, "numbers", "n", []int{}, "Numbers to be added up")
	sumCmd.Flags().StringVarP(&address, "url", "u", "", "")
}

// AddSubCommands adds the sub-commands to the provided command
func AddSubCommands(cmd *cobra.Command) {
	cmd.AddCommand(sumCmd)
}
