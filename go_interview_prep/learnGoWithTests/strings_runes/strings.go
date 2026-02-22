package main

import (
	"fmt"
	"strconv"
)

func main() {
	message := "Hello, playground"
	rawMessage := `Hello\nGo` // used for raw string literals
	fmt.Println(message, rawMessage)
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(longestPalindrome("babad"))
}

func lengthOfLongestSubstring(s string) int {
	lastSeen := make(map[rune]int)

	start := 0
	maxLength := 0

	for i, char := range s {
		if pos, ok := lastSeen[char]; ok && pos >= start {
			start = pos + 1
		}
		lastSeen[char] = i
		maxLength = max(maxLength, i-start+1)
	}

	return maxLength
}

func longestPalindrome(s string) string {
	if len(s) < 1 {
		return ""
	}

	start, end := 0, 0
	for i := range s {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		length := max(len1, len2)
		if length > end-start+1 {
			start = i - (length-1)/2
			end = i + length/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {

	//	fmt.Println(left, right, s[left], s[right])
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

func reverseMap[M ~map[K]V, K comparable, V comparable](m M) map[V]K {
	n := make(map[V]K, len(m)) // Pre-allocate the new map with the same capacity
	for k, v := range m {
		n[v] = k
	}
	return n
}

func isPalindrome(x int) bool {
	return checkPalindrome(x)
}

func checkPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	s := strconv.Itoa(x)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

var romanValues = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	if s == "" {
		return 0
	}
	return convertRomanToInt(s)
}

func convertRomanToInt(s string) int {
	sum := 0
	last := 0

	for i := len(s) - 1; i >= 0; i-- {
		v, ok := romanValues[rune(s[i])]
		if !ok {
			return 0
		}

		if v < last {
			sum -= v
		} else {
			sum += v
			last = v
		}
	}

	return sum
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// Find the shortest string; the common prefix can't be longer than this.
	shortest := strs[0]
	for _, s := range strs[1:] {
		if len(s) < len(shortest) {
			shortest = s
		}
	}

	others := filterOutString(strs, shortest)
	prefix := shortest

	for len(prefix) > 0 {
		allMatch := true
		for _, s := range others {
			if len(s) < len(prefix) || s[:len(prefix)] != prefix {
				allMatch = false
				break
			}
		}
		if allMatch {
			return prefix
		}
		prefix = prefix[:len(prefix)-1]
	}

	return ""
}

func filterOutString(slice []string, removeStr string) []string {
	newSlice := make([]string, 0, len(slice))
	removed := false

	for _, item := range slice {
		if !removed && item == removeStr {
			removed = true
			continue
		}
		newSlice = append(newSlice, item)
	}

	return newSlice
}
