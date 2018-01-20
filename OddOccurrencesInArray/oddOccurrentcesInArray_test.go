package oddOccurrencesInArray

import "testing"

func TestSolution(t *testing.T) {
	testCases := []struct {
		description string
		input       []int
		expect      int
	}{
		{
			description: "testing 1",
			input:       []int{9, 3, 9, 3, 9, 7, 9},
			expect:      7,
		},
	}

	for _, tt := range testCases {
		actual := Solution(tt.input)
		if tt.expect != actual {
			t.Errorf("%s: expect:%v != actual:%v", tt.description, tt.expect, actual)
		}
	}
}
