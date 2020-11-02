// Given a string, you need to reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.

// Example 1:
// Input: "Let's take LeetCode contest"
// Output: "s'teL ekat edoCteeL tsetnoc"
// Note: In the string, each word is separated by single space and there will not be any extra space in the string.

package main

func reverseWords(s string) string {
	ret := []byte(s)
	for i, j := 0, 0; j < len(ret); i, j = j+1, j+1 {
		for ; j < len(ret) && ret[j] != ' '; j++ {
		}
		for k, l := i, j-1; k < l; k, l = k+1, l-1 {
			ret[k], ret[l] = ret[l], ret[k]
		}
	}
	return string(ret)
}
