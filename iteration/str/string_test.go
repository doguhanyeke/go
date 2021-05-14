package str

import "testing"

func TestStr(t *testing.T) {
	got := CompareStrings("a", "b")
	expected := -1
	if got != expected {
		t.Errorf("got %d expected %d", got, expected)
	}
}
