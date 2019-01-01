package main

// Sum accepts an array of numbers and returns the sum of its elements
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

// SumAll accepts any number of slices and retunrs a slice whose elements are the sum of input slices
func SumAll(slicesToSum ...[]int) []int {
	var sum []int

	for _, slice := range slicesToSum {
		sum = append(sum, Sum(slice))
	}

	return sum
}

// SumAllTails accepts any number of slices and returns a slice whose elements are the sum of tails of input slices
func SumAllTails(slicesToSum ...[]int) []int {
	var sum []int

	for _, slice := range slicesToSum {
		tail := slice[1:]
		sum = append(sum, Sum(tail))
	}

	return sum
}

func main() {
}
