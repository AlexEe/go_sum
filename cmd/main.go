package main

import (
	cli "goSum/cli"
)

var numbers []int32

func init() {
	// get cli tool flags with numbers to sum up

}

func main() {
	cli.Execute()
	// numbers = []int32{1, 2, 3}
	// result, err := sum.Calculate(numbers)
	// if err != nil {
	// 	os.Exit(1)
	// }
	// fmt.Println(result)
}
