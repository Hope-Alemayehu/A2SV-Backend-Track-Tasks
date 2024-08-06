package main

import "testing"

func TestCalculateAverage(t *testing.T) {
	var tests = []struct {
		totalInput       float32
		numberOfSubjects int
		expectedAverage  float32
	}{
		{564, 6, 94},
		{100, 10, 10},
		{234.33, 3, 78.11},
	}

	for _, test := range tests {
		if output := calculateAverage(test.totalInput, test.numberOfSubjects); output != test.expectedAverage {
			t.Error("Test Failed: {} inputtedTotal, {} inputtedNumberOfSubjects, {} expected, recieved: {} ", test.totalInput, test.numberOfSubjects, test.expectedAverage, output)
		}
	}

}
