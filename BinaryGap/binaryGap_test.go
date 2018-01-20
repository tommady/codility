package binaryGap

import "testing"

func TestSolution(t *testing.T) {
	testCases := []struct {
		description string
		input       int
		expect      int
	}{
		{
			description: "testing 1",
			input:       529,
			expect:      4,
		},
		{
			description: "testing 2",
			input:       9,
			expect:      2,
		},
		{
			description: "testing 3",
			input:       20,
			expect:      1,
		},
		{
			description: "testing 4",
			input:       15,
			expect:      0,
		},
	}

	for _, tt := range testCases {
		actual := Solution(tt.input)
		if tt.expect != actual {
			t.Errorf("%s: expect:%v != actual:%v", tt.description, tt.expect, actual)
		}
	}
}
