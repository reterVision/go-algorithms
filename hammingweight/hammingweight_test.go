package hammingweight

import (
	"testing"
)

func TestHammingWeight(t *testing.T) {
	var testData = []struct {
		name   string
		input  int
		output int
	}{
		{
			name:   "Happy Path -- 8",
			input:  8,
			output: 1,
		},
		{
			name:   "Happy Path -- 7",
			input:  7,
			output: 3,
		},
		{
			name:   "Happy Path -- Float",
			input:  7.0,
			output: 3,
		},
	}

	for _, test := range testData {
		if test.output != NumberOfSetBits(test.input) {
			t.Errorf("%s Failed", test.name)
		}
	}
}
