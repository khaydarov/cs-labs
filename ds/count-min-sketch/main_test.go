package count_min_sketch

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	depth    int
	width    int
	items    []string
	expected map[string]int
}{
	{
		depth: 3,
		width: 10,
		items: []string{"hello", "world", "hello", "world", "hello"},
		expected: map[string]int{
			"hello": 3,
			"world": 2,
			"foo":   0,
		},
	},
	{
		depth: 3,
		width: 10,
		items: []string{"hello", "world", "hello", "world", "hello", "foo", "bar", "baz"},
		expected: map[string]int{
			"hello": 3,
			"world": 2,
			"foo":   1,
			"baz":   1,
		},
	},
}

func TestCountMinSketch(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			cm := NewCountMinSketch(tc.depth, tc.width)
			for _, item := range tc.items {
				cm.Insert(item)
			}

			for item, expected := range tc.expected {
				if cm.Query(item) != expected {
					t.Errorf("Expected %d for item %s, but got %d", expected, item, cm.Query(item))
				}
			}
		})
	}
}

func BenchmarkCountMinSketch(b *testing.B) {
	cm := NewCountMinSketch(3, 10)
	for i := 0; i < b.N; i++ {
		for _, item := range testCases[0].items {
			cm.Insert(item)
		}
	}
}

func BenchmarkCountMinSketchQuery(b *testing.B) {
	cm := NewCountMinSketch(3, 10)
	for i := 0; i < b.N; i++ {
		for _, item := range testCases[0].items {
			cm.Insert(item)
		}

		fmt.Println(cm.Query("hello"))
	}
}

func BenchmarkHashMapApproach(b *testing.B) {
	hashMap := make(map[string]int)
	for i := 0; i < b.N; i++ {
		for _, item := range testCases[0].items {
			hashMap[item]++
		}
	}
}

func BenchmarkHashMapApproachQuery(b *testing.B) {
	hashMap := make(map[string]int)
	for i := 0; i < b.N; i++ {
		for _, item := range testCases[0].items {
			hashMap[item]++
		}

		fmt.Println(hashMap["hello"])
	}
}
