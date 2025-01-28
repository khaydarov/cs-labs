package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	num    int
	expect string
}{
	{
		num:    26,
		expect: "1a",
	},
	{
		num:    -1,
		expect: "ffffffff",
	},
	{
		num:    0,
		expect: "0",
	},
	{
		num:    16,
		expect: "10",
	},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			answer := toHex(tc.num)
			if answer != tc.expect {
				t.Errorf("expected: %s, got: %s", tc.expect, answer)
			}
		})
	}
}
