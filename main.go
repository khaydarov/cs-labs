package main

import (
	"fmt"
	"sort"
)

type Elem struct {
	Row      int
	Soldiers int
}

func Construct(capacity int) Heap {
	return Heap{
		size:     0,
		capacity: capacity,
		data:     make([]Elem, capacity),
	}
}

type Heap struct {
	size     int
	capacity int
	data     []Elem
}

// Left Returns index of left node
func (h *Heap) Left(i int) int {
	return 2*i + 1
}

// Right Returns index of right node
func (h *Heap) Right(i int) int {
	return 2*i + 2
}

// Parent Returns index of parent node
func (h *Heap) Parent(i int) int {
	return (i - 1) / 2
}

// Insert adds new item to heap
func (h *Heap) Insert(k Elem) {
	if h.size == h.capacity {
		return
	}

	h.size++
	i := h.size - 1
	h.data[i] = k

	for i != 0 && h.data[h.Parent(i)].Soldiers < h.data[i].Soldiers {
		h.data[h.Parent(i)], h.data[i] = h.data[i], h.data[h.Parent(i)]
		i = h.Parent(i)
	}
}

// Get Returns first heap element: min or max value
func (h *Heap) Get() Elem {
	if h.size <= 0 {
		return Elem{}
	}

	return h.data[0]
}

// Heapify adjusts heap data
func (h *Heap) Heapify(i int) {
	left := h.Left(i)
	right := h.Right(i)

	biggest := i
	if left < h.size && h.data[left].Soldiers > h.data[biggest].Soldiers {
		biggest = left
	}

	if right < h.size && h.data[right].Soldiers > h.data[biggest].Soldiers {
		biggest = right
	}

	if biggest != i {
		h.data[i], h.data[biggest] = h.data[biggest], h.data[i]
		h.Heapify(biggest)
	}
}

func (h *Heap) Extract() Elem {
	if h.size <= 0 {
		return Elem{}
	}

	returnValue := h.data[0]
	h.data[0] = h.data[h.size-1]
	h.data[h.size-1] = Elem{}
	h.size--

	h.Heapify(0)

	return returnValue
}

func (h *Heap) Size() int {
	return h.size
}

// Approach 1
// map[row] = count
// m[0]=2
// m[1]=4
// m[2]=1
// m[3]=2
// m[4]=5
// create a struct of struct { Row, Soldiers }
// sort it
// return first k rows
//
// TC: N for map and struct, NlogN for sorting
// SC: N

// Approach 2
// 4 2 1 | 2 -> 2 2 1
func kWeakestRows(mat [][]int, k int) []int {
	m := make(map[int]Elem)
	for i := 0; i < len(mat); i++ {
		max := 0
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 1 {
				max++
			}
		}
		m[i] = Elem{
			Row:      i,
			Soldiers: max,
		}
	}

	rows := make([]Elem, len(mat))
	for i := 0; i < len(mat); i++ {
		rows[i] = m[i]
	}

	sort.Slice(rows, func(i, j int) bool {
		return rows[i].Soldiers < rows[j].Soldiers || (rows[i].Soldiers == rows[j].Soldiers && rows[i].Row < rows[j].Row)
	})

	var result []int
	for i := 0; i < k; i++ {
		result = append(result, rows[i].Row)
	}
	return result
}

func main() {
	// input := [][]int{
	// 	[]int{1, 1, 0, 0, 0},
	// 	[]int{1, 1, 1, 1, 0},
	// 	[]int{1, 0, 0, 0, 0},
	// 	[]int{1, 1, 0, 0, 0},
	// 	[]int{1, 1, 1, 1, 1},
	// }
	// k := 3
	input := [][]int{
		[]int{1, 0, 0, 0},
		[]int{1, 1, 1, 1},
		[]int{1, 0, 0, 0},
		[]int{1, 0, 0, 0},
	}
	k := 2
	r := kWeakestRows(input, k)
	fmt.Println(r)
}
