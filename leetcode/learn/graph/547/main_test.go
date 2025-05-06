package main

import "testing"

func Test_findCircleNum(t *testing.T) {
	tests := []struct {
		name        string
		isConnected [][]int
		want        int
	}{
		{
			name: "один город",
			isConnected: [][]int{
				{1},
			},
			want: 1,
		},
		{
			name: "два связанных города",
			isConnected: [][]int{
				{1, 1},
				{1, 1},
			},
			want: 1,
		},
		{
			name: "два несвязанных города",
			isConnected: [][]int{
				{1, 0},
				{0, 1},
			},
			want: 2,
		},
		{
			name: "три города с двумя провинциями",
			isConnected: [][]int{
				{1, 1, 0},
				{1, 1, 0},
				{0, 0, 1},
			},
			want: 2,
		},
		{
			name: "три связанных города",
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
