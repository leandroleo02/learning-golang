package bufferr

import (
	"testing"
)

func TestAnagramDifferentSize(t *testing.T) {
	s1 := "listen"
	s2 := "list"

	r := Anagram(s1, s2)
	if r {
		t.Fatalf("%q and %q area not anagrams", s1, s2)
	}
}

func TestAnagramDifferentStringSameSize(t *testing.T) {
	s1 := "listen"
	s2 := "listte"

	r := Anagram(s1, s2)
	if r {
		t.Fatalf("%q and %q area not anagrams", s1, s2)
	}
}

func TestAnagramEqualStrings(t *testing.T) {
	s1 := "listen"
	s2 := "listen"

	r := Anagram(s1, s2)
	if !r {
		t.Fatalf("%q and %q area not anagrams", s1, s2)
	}
}

func TestAnagram(t *testing.T) {
	s1 := "listen"
	s2 := "silent"

	r := Anagram(s1, s2)
	if !r {
		t.Fatalf("%q and %q area not anagrams", s1, s2)
	}
}
