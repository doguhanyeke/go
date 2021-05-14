package iteration

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T) {
	got := Repeat("a", 5)
	want := "aaaaa"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	i := 0
	for i <= b.N {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	got := Repeat("b", 4)
	fmt.Println(got)
	// Output: bbbb
}
