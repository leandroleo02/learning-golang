package bufferr

// Anagram verify whether two strings area anagrams
func Anagram(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	if s1 == s2 {
		return true
	}

	lettersCount := make(map[string]int)
	for i := 0; i < len(s1); i++ {
		lettersCount[string(s1[i])]++
		lettersCount[string(s2[i])]--
	}

	for _, c := range lettersCount {
		if c != 0 {
			return false
		}
	}
	return true
}
