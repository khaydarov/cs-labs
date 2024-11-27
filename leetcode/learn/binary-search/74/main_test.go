package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	matrix [][]int
	target int
	expect bool
}{
	{
		matrix: [][]int{
			{1},
		},
		target: 0,
		expect: false,
	},
	{
		matrix: [][]int{
			{1},
		},
		target: 1,
		expect: true,
	},
	{
		matrix: [][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		},
		target: 3,
		expect: true,
	},
	{
		matrix: [][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		},
		target: 5,
		expect: true,
	},
	{
		matrix: [][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		},
		target: 34,
		expect: true,
	},
	{
		matrix: [][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		},
		target: 23,
		expect: true,
	},
	{
		matrix: [][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		},
		target: 35,
		expect: false,
	},
}

func TestSolution1(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := searchMatrix(testCase.matrix, testCase.target)
			if output != testCase.expect {
				t.Errorf("Wrong answer. Expected %v, but got %v", testCase.expect, output)
			}
		})
	}
}

func TestSolution2(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			output := searchMatrix1(testCase.matrix, testCase.target)
			if output != testCase.expect {
				t.Errorf("Wrong answer. Expected %v, but got %v", testCase.expect, output)
			}
		})
	}
}
