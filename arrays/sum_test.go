package main

import (
	"fmt"
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
