package basename

import "testing"

func TestBasename(t *testing.T) {
	s := "a/b/basename.go"
	got := basename(s)
	want := "basename"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
