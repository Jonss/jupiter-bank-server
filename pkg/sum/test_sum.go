package sum

import "testing"

// TestSum is temporary. It'll be removed when actual tests exists
func TestSum(t *testing.T) {
	got := Sum(2, 2)
	want := 4
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
