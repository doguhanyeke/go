package maps

import "testing"

func TestSearch(t *testing.T) {
	value := "this is just a test"
	dictionary := map[string]string{"test": value}

	key := "test"
	got := Search(dictionary, key)
	want := value
	if got != want {
		t.Errorf("got %s want %s, given key as %s", got, want, key)
	}
}
