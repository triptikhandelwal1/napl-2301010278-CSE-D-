package strop

import "strings"

func CountVowels(s string) int {
	count := 0

	for _, ch := range strings.ToLower(s) {
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
			count++
		}
	}

	return count
}
