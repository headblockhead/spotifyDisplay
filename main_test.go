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

func TestPad(t *testing.T) {
	tests := []struct {
		input    string
		length   int
		expected string
	}{
		{
			input:    "abc",
			length:   3,
			expected: "abc",
		},
		{
			input:    "abc",
			length:   10,
			expected: "abc       ",
		},
		{
			input:    "  123 ",
			length:   10,
			expected: "  123     ",
		},
		{
			input:    "abc",
			length:   2,
			expected: "abc",
		},
	}
	for _, test := range tests {
		actual := pad(test.input, test.length)
		if actual != test.expected {
			t.Errorf("for input %q to=%v, expected %q, got %q", test.input, test.length, test.expected, actual)
		}
	}
}
