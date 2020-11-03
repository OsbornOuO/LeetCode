// Given two strings s and t , write a function to determine if t is an anagram of s.

package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	result := make([]int, 26)
	for _, ss := range s {
		result[ss-'a']++
	}

	for _, tt := range t {
		result[tt-'a']--
		if result[tt-'a'] < 0 {
			return false
		}
	}
	return true
}
