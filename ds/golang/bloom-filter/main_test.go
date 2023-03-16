package main

import (
	"fmt"
	"testing"
)

var words = []string{
	"hello",
	"bloom",
	"bloom-filter",
	"filter",
	"word",
	"golang",
}

var testCases = []struct {
	item   string
	expect bool
}{
	{
		"hello",
		true,
	},
	{
		"bloom",
		true,
	},
	{
		"bloom-filter",
		true,
	},
	{
		"filter",
		true,
	},
	{
		"word",
		true,
	},
	{
		"golang",
		true,
	},
	{
		"non-existed",
		false,
	},
	{
		"some-key",
		false,
	},
}

func TestSolution(t *testing.T) {
	bloomFilter := NewBloomFilter(1000, 0.05)
	for _, v := range words {
		bloomFilter.Insert(v)
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := bloomFilter.Check(testCase.item)

			if output != testCase.expect {
				t.Errorf("Wrong answer for %v", testCase.item)
			}
		})
	}
}
