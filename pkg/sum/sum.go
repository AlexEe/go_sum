package sum

// Calculate creates sum of all numbers passed into function
func Calculate(numbers []int32) (int32, error) {
	var result int32
	result = 0
	for _, number := range numbers {
		result = result + number
	}
	return result, nil
}
