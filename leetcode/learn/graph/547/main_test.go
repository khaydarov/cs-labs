package main

import "testing"

func Test_findCircleNum(t *testing.T) {
	tests := []struct {
		name        string
		isConnected [][]int
		want        int
	}{
		{
			name: "one city",
			isConnected: [][]int{
				{1},
			},
			want: 1,
		},
		{
			name: "two connected cities",
			isConnected: [][]int{
				{1, 1},
				{1, 1},
			},
			want: 1,
		},
		{
			name: "two non-connected cities",
			isConnected: [][]int{
				{1, 0},
				{0, 1},
			},
			want: 2,
		},
		{
			name: "three cities with two provinces",
			isConnected: [][]int{
				{1, 1, 0},
				{1, 1, 0},
				{0, 0, 1},
			},
			want: 2,
		},
		{
			name: "three connected cities",
			isConnected: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCircleNum(tt.isConnected); got != tt.want {
				t.Errorf("findCircleNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
