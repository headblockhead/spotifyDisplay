package main

import "testing"

func TestSubstring(t *testing.T) {
	tests := []struct {
		input    string
		from     int
		length   int
		expected string
	}{
		{
			input:    "abcdef",
			from:     1,
			length:   3,
			expected: "bcd",
		},
		{
			input:    "abcdef",
			from:     0,
			length:   10,
			expected: "abcdef",
		},
		{
			input:    "abcdef",
			from:     7,
			length:   3,
			expected: "",
		},
	}
	for _, test := range tests {
		actual := substring(test.input, test.from, test.length)
		if actual != test.expected {
			t.Errorf("for input %q from=%v to=%v, expected %q, got %q", test.input, test.from, test.length, test.expected, actual)
		}
	}
}
