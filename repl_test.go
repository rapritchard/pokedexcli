package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{
			input: "hello world",
			expected: []string{
				"hello",
				"world",
			},
		},
		{
			input: "hElLo WoRlD",
			expected: []string{
				"hello",
				"world",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			actual := cleanInput(tc.input)
			if len(actual) != len(tc.expected) {
				t.Errorf("the lengths are not equal: %v vs %v", len(actual), len(tc.expected))
			}
			for i := range actual {
				actualWord := actual[i]
				expectedWord := tc.expected[i]
				if actualWord != expectedWord {
					t.Errorf("%v does not equal %v", actualWord, expectedWord)
				}
			}
		})
	}
}
