package main

import "testing"

func TestLFUCache(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
		ops      []string
		inputs   [][]int
		want     []int
	}{
		{
			name:     "basic operations",
			capacity: 2,
			ops:      []string{"put", "put", "get", "put", "get", "get"},
			inputs:   [][]int{{1, 1}, {2, 2}, {1}, {3, 3}, {2}, {3}},
			want:     []int{-1, -1, 1, -1, -1, 3},
		},
		{
			name:     "zero capacity",
			capacity: 0,
			ops:      []string{"put", "get"},
			inputs:   [][]int{{1, 1}, {1}},
			want:     []int{-1, -1},
		},
		{
			name:     "update existing value",
			capacity: 2,
			ops:      []string{"put", "put", "get", "put", "get"},
			inputs:   [][]int{{1, 1}, {1, 2}, {1}, {3, 3}, {1}},
			want:     []int{-1, -1, 2, -1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := NewLFUCache(tt.capacity)
			var got []int

			for i, op := range tt.ops {
				switch op {
				case "put":
					cache.Put(tt.inputs[i][0], tt.inputs[i][1])
					got = append(got, -1)
				case "get":
					val := cache.Get(tt.inputs[i][0])
					got = append(got, val)
				}
			}

			for i, want := range tt.want {
				if got[i] != want {
					t.Errorf("operation %d: got %d, want %d", i, got[i], want)
				}
			}
		})
	}
}
