package lc125validpalindrome

import (
	"regexp"
	"strings"
)

func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	s = regexp.MustCompile(`[^A-Za-z0-9]`).ReplaceAllString(s, "")

	if len(strings.TrimSpace(s)) <= 1 {
		return true
	}

	left := 0
	right := len(s) - 1

	for left < right {
		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}
	return true
}
