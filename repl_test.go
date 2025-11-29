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
		input: "gotta catch 'em all",
		expected: []string{"gotta", "catch", "'em", "all"},
	},
	{
		input: "Charmander Bulbasaur PIKACHU",
		expected: []string{"charmander", "bulbasaur", "pikachu"},
	},
	{
		input: "shabingus",
		expected: []string{"shabingus"},
	},
}

for _, c := range cases {
	actual := cleanInput(c.input)
	if len(actual) != len(c.expected) {
		t.Errorf("length of slice not equal to expected slice")
		t.Fail()
	}

	for i := range actual {
		word := actual[i]
		expectedWord := c.expected[i]
		if word != expectedWord {
			t.Errorf("actual word does not match expected word")
			t.Fail()
		}
	}
}
}
