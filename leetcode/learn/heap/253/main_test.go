package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	intervals [][]int
	answer    int
}{
	{
		intervals: [][]int{
			{0, 30},
			{5, 10},
			{15, 20},
		},
		answer: 2,
	},
	{
		intervals: [][]int{
			{0, 30},
			{5, 10},
			{0, 20},
		},
		answer: 3,
	},
	{
		intervals: [][]int{
			{7, 10},
			{2, 4},
		},
		answer: 1,
	},
	{
		intervals: [][]int{
			{1, 10},
			{5, 12},
			{6, 8},
			{1, 5},
			{5, 10},
		},
		answer: 4,
	},
}

func TestSolution(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			result := minMeetingRooms(testCase.intervals)
			if result != testCase.answer {
				t.Errorf("Wrong answer. Expected %v, but got %v", testCase.answer, result)
			}
		})
	}
}
