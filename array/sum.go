package array

// Sum will take an array of numbers and return the total
func Sum(numbers [5]int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}
