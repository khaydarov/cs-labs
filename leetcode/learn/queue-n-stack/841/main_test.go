package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	rooms  [][]int
	expect bool
}{
	{
		[][]int{
			{1},
			{2},
			{3},
			{},
		},
		true,
	},
	{
		[][]int{
			{1, 3},
			{3, 0, 1},
			{2},
			{0},
		},
		false,
	},
	{
		[][]int{
			{1},
			{0},
			{0},
		},
		false,
	},
	{
		[][]int{
			{1},
			{1},
			{2},
		},
		false,
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := canVisitAllRooms(testCase.rooms)
			if output != testCase.expect {
				t.Errorf("Wrong answer. Expected %v but got %v", testCase.expect, output)
			}
		})
	}
}
