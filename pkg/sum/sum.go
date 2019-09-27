package sum

// Sum creates sum of all numbers passed into function
func Sum(numbers []int) (int, error) {
	result := 0
	for _, number := range numbers {
		result = result + number
	}
	return result, nil
}
