package main

import "testing"

var testCases = []struct {
	seats [][]byte
	want  int
}{
	{
		seats: [][]byte{
			{'#', '.', '#', '#', '.', '#'},
			{'.', '#', '#', '#', '#', '.'},
			{'#', '.', '#', '#', '.', '#'},
		},
		want: 4,
	},
	{
		seats: [][]byte{
			{'.', '#'},
			{'#', '#'},
			{'#', '.'},
		},
		want: 2,
	},
	{
		seats: [][]byte{
			{'#', '.', '.', '.', '#'},
			{'.', '#', '.', '#', '.'},
			{'.', '.', '#', '.', '.'},
			{'.', '#', '.', '#', '.'},
			{'#', '.', '.', '.', '#'},
		},
		want: 10,
	},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		got := maxStudents(tc.seats)
		if got != tc.want {
			t.Errorf("TestSolution(%d): got: %v, want: %v", i, got, tc.want)
		}
	}
}
