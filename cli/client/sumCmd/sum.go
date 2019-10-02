package sumCmd

import (
	"goSum/pkg/client"

	"github.com/spf13/cobra"
)

var (
	numbers []int
	address string
)

// rootCmd represents the base command when called without any subcommands
var sumCmd = &cobra.Command{
	Use:     "sum",
	Short:   "Add numbers entered on the Command Line",
	Example: "sum -n 4,1,-2 -u localhost:8080",
	Run: func(cmd *cobra.Command, args []string) {
		client.Connect(numbers, address)
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
