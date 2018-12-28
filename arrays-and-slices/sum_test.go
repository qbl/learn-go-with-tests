package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("Sum of five numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		actual := Sum(numbers)
		expected := 15

		if actual != expected {
			t.Errorf("actual: %d; expected: %d", actual, expected)
		}
	})

	t.Run("Sum of three numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		actual := Sum(numbers)
		expected := 15

		if actual != expected {
			t.Errorf("actual: %d; expected: %d", actual, expected)
		}
	})
}
