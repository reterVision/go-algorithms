package kmp

import (
	"testing"
)

func TestKMP(t *testing.T) {
	testScenarios := []struct {
		source string
		target string
		index  int
	}{
		{
			source: "mississipi",
			target: "issip",
			index:  4,
		},
		{
			source: "mississipi",
			target: "a",
			index:  -1,
		},
		{
			source: "hello,world",
			target: "ll",
			index:  2,
		},
		{
			source: "banananobano",
			target: "nano",
			index:  4,
		},
	}

	for _, s := range testScenarios {
		index := KMP(s.source, s.target)
		if index != s.index {
			t.Errorf("source: %s, target: %s, expected %d but get %d", s.source, s.target, s.index, index)
		}
	}
}
