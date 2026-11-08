package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {

	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " HEllo world ",
			expected: []string{"hello", "world"},
		},
	}

	for i := range cases {
		actual := cleanInput(cases[i].input)

		if len(actual) != len(cases[i].expected) {
			t.Errorf("case %d: expected %v, got %v", i, cases[i].expected, actual)
		}
		for j := range actual {
			word := actual[j]
			expextedWord := cases[i].expected[j]
			if word != expextedWord {
				t.Errorf("case %d: expected %v, got %v", i, cases[i].expected, actual)
			}
		}
	}

}
