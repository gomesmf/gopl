package main

import "testing"

func TestIsAnagram(t *testing.T) {
	t.Run("is anagram", func(t *testing.T) {
		s1 := "matheus"
		s2 := "suehtam"
		got := isAnagram(s1, s2)
		want := true
		if got != want {
			t.Errorf("isAnagram(%q, %q) = %v, want %v", s1, s2, got, want)
		}
	})

	t.Run("is not anagram", func(t *testing.T) {
		s1 := "matheus"
		s2 := "gomes"
		got := isAnagram(s1, s2)
		want := false
		if got != want {
			t.Errorf("isAnagram(%q, %q) = %v, want %v", s1, s2, got, want)
		}
	})
}
