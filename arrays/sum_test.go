package main

import (
	"fmt"
	"slices"
	"testing"
)

func ExampleSum() {
	sum_array := Sum([]int{1, 2, 3, 4, 5})
	fmt.Println(sum_array)
	// Output: 15
}

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given %v", got, want, numbers)
	}
}

func ExampleSumAll() {
	sum_array := SumAll([]int{1, 2}, []int{4, 5})
	fmt.Println(sum_array)
	// Output: [3 9]
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
