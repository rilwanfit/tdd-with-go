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
func SumAll(numbersTosum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersTosum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

// SumAllTails which calculates the totals of the "tails" of each slice. The tail of a collection is all the items apart from the first one (the "head")
func SumAllTails(numbersTosum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersTosum {
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}

	return sums
}
