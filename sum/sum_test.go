package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Sum of 5 elements in array", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		if got != want {
			t.Errorf("got %d want %d, given %v", got, want, numbers)
		}
	})
	// t.Run("Sum of dynamic sizes", func(t *testing.T) {
	// 	numbers := []int{1, 2, 3}
	// 	got := Sum(numbers)
	// 	want := 6
	// 	if got != want {
	// 		t.Errorf("got %d want %d, given %v", got, want, numbers)
	// 	}
	// })
}

func TestSumAll(t *testing.T) {
	t.Run("SumAll", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3, 4}, []int{3, 4, 5})
		want := []int{10, 12}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got []int, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("Tail", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{3, 4, 5})
		want := []int{5, 9}
		checkSums(t, got, want)
	})

	t.Run("Safely Sum Empty Slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}
