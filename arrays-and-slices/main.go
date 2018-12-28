package main

import "fmt"

// Sum accepts an array of numbers and return the sum of its elements
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(Sum(numbers))
}
