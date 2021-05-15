package sum

import "testing"

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
