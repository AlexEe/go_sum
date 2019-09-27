package main

import (
	"fmt"
	"goSum/pkg/sum"
	"os"
)

var numbers []int

func init() {
	// get cli tool flags with numbers to sum up

}

func main() {
	numbers = []int{1, 2, 3}
	result, err := sum.Sum(numbers)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(result)
}
