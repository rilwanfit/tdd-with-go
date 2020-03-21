package array

// Sum will take an array of numbers and return the total
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

// SumAll which will take a varying number of slices, returning a new slice containing the totals for each slice passed in.
func SumAll(number1, number2 []int) {

}
