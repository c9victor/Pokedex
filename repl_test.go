package main

import "testing"

// testing documentation: https://pkg.go.dev/testing

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
			input:    " ",
			expected: []string{""},
		},
		{
			input:    "   Hello dear frIends",
			expected: []string{"hello", "dear", "friends"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Fail: \nactual list: %v\nexpected list: %v\nactual length: %d\nexpected length: %d\n\n",
				actual, c.expected, len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Fail:\nactual word: %v\nexpected word: %v\n\n", word, expectedWord)
			}
		}
	}
}
