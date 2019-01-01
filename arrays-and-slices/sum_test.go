package main

import (
	"reflect"
	"testing"
)

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

func TestSumAll(t *testing.T) {
	actual := SumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual: %d; expected: %d", actual, expected)
	}
}

func TestSumAllTails(t *testing.T) {
	t.Run("Sum all tails of non-empty slices", func(t *testing.T) {
		actual := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("actual: %d; expected: %d", actual, expected)
		}
	})

	t.Run("Sum all tails of empty slices", func(t *testing.T) {
		actual := SumAllTails([]int{}, []int{0, 9})
		expected := []int{0, 9}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("actual: %d; expected: %d", actual, expected)
		}
	})
}
