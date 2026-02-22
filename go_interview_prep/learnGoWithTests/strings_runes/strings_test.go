package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{" ", 1},
		{"au", 2},
		{"dvdf", 3},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring(%q) = %d; want %d", tt.s, got, tt.want)
			}
		})
	}
}

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"babad", "bab"}, // "aba" is also valid, but our algo picks "bab"
		{"cbbd", "bb"},
		{"a", "a"},
		{"ac", "a"},
		{"racecar", "racecar"},
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := longestPalindrome(tt.s); got != tt.want {
				// Special case for babad where aba is also valid
				if tt.s == "babad" && got == "aba" {
					return
				}
				t.Errorf("longestPalindrome(%q) = %q; want %q", tt.s, got, tt.want)
			}
		})
	}
}
func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		name     string
		input    int
		expected bool
	}{
		{name: "121", input: 121, expected: true},
		{name: "10", input: 10, expected: false},
		{name: "-121", input: -121, expected: false},
	}

	for _, tt := range cases {
		tt := tt // rebinding avoids accidental capture issues (e.g., if t.Parallel() is added later)
		t.Run(tt.name, func(t *testing.T) {
			got := checkPalindrome(tt.input)
			if got != tt.expected {
				t.Errorf("checkPalindrome(%d) = %v; want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  string
	}{
		{
			name:  "empty slice",
			input: []string{},
			want:  "",
		},
		{
			name:  "single string",
			input: []string{"alone"},
			want:  "alone",
		},
		{
			name:  "typical case",
			input: []string{"flower", "flow", "flight"},
			want:  "fl",
		},
		{
			name:  "no common prefix",
			input: []string{"dog", "racecar", "car"},
			want:  "",
		},
		{
			name:  "includes empty string",
			input: []string{"", "abc", "ab"},
			want:  "",
		},
		{
			name:  "all identical",
			input: []string{"same", "same", "same"},
			want:  "same",
		},
		{
			name:  "shortest string is the prefix of others",
			input: []string{"pre", "prefix", "prevent"},
			want:  "pre",
		},
		{
			name:  "shortest appears more than once",
			input: []string{"ab", "ab", "abc"},
			want:  "ab",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := longestCommonPrefix(tc.input)
			if got != tc.want {
				t.Fatalf("longestCommonPrefix(%v) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}
