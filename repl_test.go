package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input  string
		output []string
	}{
		{
			input:  "  hello world  ",
			output: []string{"hello", "world"},
		},
		{
			input:  "         ",
			output: []string{},
		},
		{
			input:  "		hdjsd 		kndknv",
			output: []string{"hdjsd", "kndknv"},
		},
		{
			input:  "Meow,  Meow",
			output: []string{"meow,", "meow"},
		},
	}

	for k, cs := range cases {
		actualOutput := cleanInput(cs.input)

		if len(actualOutput) != len(cs.output) {
			t.Errorf("the lengths of the %dth case are not equal :-  actual: %v != expected %v", k, len(actualOutput), len(cs.output))
			continue
		} else {
			for i := range actualOutput {
				if actualOutput[i] != cs.output[i] {
					t.Errorf("the values of the ith case are not equal :-  actual: %v != expected %v", actualOutput[i], cs.output[i])
					break
				}
			}
		}
	}
}
