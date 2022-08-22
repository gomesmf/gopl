package comma

import "testing"

func TestComma(t *testing.T) {
	s := "12345"
	got := comma(s)
	want := "12,345"
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}

	s = "12345678"
	got = comma(s)
	want = "12,345,678"
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
