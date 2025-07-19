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
		input:    "  hello  world  ",
		expected: []string{"hello", "world"},
	},
	{
		input:    "hello",
		expected: []string{"hello"},
	},
	{
		input:    "hi  hello  world  ",
		expected: []string{"hi", "hello", "world"},
	},
}

for _, c := range cases {
	result := cleanInput(c.input)
	if len(result) != len(c.expected) {
		t.Errorf("lengths do not match '%v' vs '%v'", result, c.expected)
		continue
	}


	for i := range result {
		word := result[i]
		expectedWord := c.expected[i]
		if word != expectedWord {
			t.Errorf("cleanInput(%v) == %v, expected %v", c.input, result, c.expected)
		}
	}
}
}
