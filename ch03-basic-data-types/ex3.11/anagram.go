package main

// isAnagram reports whether two strings are anagrams of each other.
func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	m1 := make(map[rune]int)
	for _, c := range s1 {
		m1[c]++
	}

	m2 := make(map[rune]int)
	for _, c := range s2 {
		m2[c]++
	}

	for c1, v1 := range m1 {
		v2, ok := m2[c1]
		if !ok {
			return false
		}
		if v1 != v2 {
			return false
		}
	}
	return true
}
