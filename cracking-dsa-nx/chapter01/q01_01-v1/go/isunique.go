package isunique

type isUniqueImpl func(string) bool

// BruteForce checks if a string has all unique characters using nested loops.
func BruteForce(s string) bool {
	runes := []rune(s)
	for i := 0; i < len(runes)-1; i++ {
		for j := i + 1; j < len(runes); j++ {
			if runes[i] == runes[j] {
				return false
			}
		}
	}
	return true
}

func WithSet(s string) bool {
	seen := map[rune]bool{}
	for _, r := range s {
		if _, exists := seen[r]; exists {
			return false
		}
		seen[r] = true
	}
	return true
}
