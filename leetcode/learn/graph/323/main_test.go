package main

import "testing"

func TestCountComponents(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		edges    [][]int
		expected int
	}{
		{
			name:     "empty graph",
			n:        5,
			edges:    [][]int{},
			expected: 5,
		},
		{
			name:     "one component",
			n:        5,
			edges:    [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}},
			expected: 1,
		},
		{
			name:     "two components",
			n:        5,
			edges:    [][]int{{0, 1}, {1, 2}, {3, 4}},
			expected: 2,
		},
		{
			name:     "three components",
			n:        6,
			edges:    [][]int{{0, 1}, {2, 3}, {4, 5}},
			expected: 3,
		},
		{
			name:     "graph with cycle",
			n:        4,
			edges:    [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countComponents(tt.n, tt.edges)
			if result != tt.expected {
				t.Errorf("countComponents(%d, %v) = %d; expected %d",
					tt.n, tt.edges, result, tt.expected)
			}
		})
	}
}
