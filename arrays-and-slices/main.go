package main

import "fmt"

// Sum accepts an array of numbers and return the sum of its elements
func Sum(numbers [5]int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}

	return sum
}

func main() {
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println(Sum(numbers))
}
