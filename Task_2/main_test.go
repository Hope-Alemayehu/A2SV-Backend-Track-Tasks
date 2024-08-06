package main

import (
	"testing"
)

// areMapsEqual checks if two maps are equal.
func areMapsEqual(m1, m2 map[string]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}
	return true
}

// TestWordFrequencyCounter tests the wordFrequencyCounter function.
func TestWordFrequencyCounter(t *testing.T) {
	tests := []struct {
		input        string
		expectedFreq map[string]int
	}{
		{"we aRe. frw We", map[string]int{"we": 2, "are": 1, "frw": 1}},
		{"Hello, world! Hello again.", map[string]int{"hello": 2, "world": 1, "again": 1}},
		{"", map[string]int{}},
		{"!!!", map[string]int{}},
		{"Test, test; test: TEST!", map[string]int{"test": 4}},
		{"This is a test. This test is simple.", map[string]int{"this": 2, "is": 2, "a": 1, "test": 2, "simple": 1}},
		{"Numbers 1234 and symbols #$%^&*()", map[string]int{"numbers": 1, "1234": 1, "and": 1, "symbols": 1}},
		{"Mixed CASE words. MiXeD caSe.", map[string]int{"mixed": 2, "case": 2, "words": 1}},
		{"Special characters: ~`!@#$%^&*()_+-=[]{}|;:',<.>/?", map[string]int{"special": 1, "characters": 1}},
		{"Newline\nand\ttabs", map[string]int{"newline": 1, "and": 1, "tabs": 1}},
	}

	for _, test := range tests {
		output := wordFrequencyCounter(test.input)
		if !areMapsEqual(output, test.expectedFreq) {
			t.Errorf("For input '%s', expected %v, but got %v", test.input, test.expectedFreq, output)
		}
	}
}

// TestCheckPalindrome tests the checkPalindrome function.
func TestCheckPalindrome(t *testing.T) {
	tests := []struct {
		inputString    string
		expectedAnswer bool
	}{
		{"EREweq", false},
		{"erwwre", true},
		{"we", false},
		{"A man a plan a canal Panama", true}, // Adding a case-insensitive palindrome
		{"", true},                            // An empty string is considered a palindrome
		{"12321", true},                       // A numeric palindrome
	}

	for _, test := range tests {
		if output := checkPalindrome(test.inputString); output != test.expectedAnswer {
			t.Errorf("For input '%s', expected %v, but got %v", test.inputString, test.expectedAnswer, output)
		}
	}
}
