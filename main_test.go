package main

import "testing"

func TestValidateInput(t *testing.T) {

	var validArray = []struct {
		input    string
		expected int
	}{
		{"23", 23},
		{"42", 42},
		{"16", 16},
		{"64", 64},
	}

	for _, test := range validArray {
		if output, err := validateInput(test.input); output != test.expected {
			t.Error("Test Failed: {}inputted, {}expected, Received: {}", test.input, test.expected, err)
		}
	}
}

func TestJosephusSimulation(t *testing.T) {

	var josephusSimulationArray = []struct {
		input    int
		expected int
	}{
		{41, 19},
		{64, 1},
		{34, 5},
		{144, 33},
	}

	for _, test := range josephusSimulationArray {
		if output := josephusSimulation(test.input); output != test.expected {
			t.Error("Test Failed: {}inputted, {}expected, Received: {}", test.input, test.expected, output)
		}
	}
}
