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
	length := len(slicesToSum)
	sum := make([]int, length)

	for i, slice := range slicesToSum {
		sum[i] = Sum(slice)
	}

	return sum
}

func main() {
}
