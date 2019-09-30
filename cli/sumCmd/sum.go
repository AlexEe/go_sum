package sumCmd

import (
	"goSum/pkg/sum"

	"github.com/spf13/cobra"
)

var numbers []int32

var sumCmd = &cobra.Command{
	Use:     "sum",
	Short:   "Calculate the sum of all numbers entered on CLI",
	Example: "sum -n 1, 2, 3",
	Run: func(cmd *cobra.Command, args []string) {
		sum.Add(numbers)
	},
}

func init() {
	sumCmd.Flags().Int32SliceVarP(&numbers, "numbers", "n", []int32{}, "Numbers to be added up")
}

// AddSubCommands adds the sub-commands to the provided command
func AddSubCommands(cmd *cobra.Command) {
	cmd.AddCommand(sumCmd)
}