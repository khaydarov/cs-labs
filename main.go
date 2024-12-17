package main

import (
	"fmt"
	"math"
)

type Point struct {
	Coords []int
	Dist   float64
}

func Construct(capacity int) Heap {
	return Heap{
		size:     0,
		capacity: capacity,
		data:     make([]Point, capacity),
	}
}

type Heap struct {
	size     int
	capacity int
	data     []Point
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
func (h *Heap) Insert(k Point) {
	if h.size == h.capacity {
		return
	}

	h.size++
	i := h.size - 1
	h.data[i] = k

	for i != 0 && h.data[h.Parent(i)].Dist < h.data[i].Dist {
		h.data[h.Parent(i)], h.data[i] = h.data[i], h.data[h.Parent(i)]
		i = h.Parent(i)
	}
}

// Get Returns first heap element: min or max value
func (h *Heap) Get() Point {
	if h.size <= 0 {
		return Point{}
	}

	return h.data[0]
}

// Heapify adjusts heap data
func (h *Heap) Heapify(i int) {
	left := h.Left(i)
	right := h.Right(i)

	biggest := i
	if left < h.size && h.data[left].Dist > h.data[biggest].Dist {
		biggest = left
	}

	if right < h.size && h.data[right].Dist > h.data[biggest].Dist {
		biggest = right
	}

	if biggest != i {
		h.data[i], h.data[biggest] = h.data[biggest], h.data[i]
		h.Heapify(biggest)
	}
}

func (h *Heap) Extract() Point {
	if h.size <= 0 {
		return Point{}
	}

	returnValue := h.data[0]
	h.data[0] = h.data[h.size-1]
	h.data[h.size-1] = Point{}
	h.size--

	h.Heapify(0)

	return returnValue
}

func (h *Heap) Size() int {
	return h.size
}

func calculateDistance(p []int) float64 {
	return math.Sqrt(float64(p[0]*p[0] + p[1]*p[1]))
}

func kClosest(points [][]int, k int) [][]int {
	if len(points) == k {
		return points
	}

	heap := Construct(k)
	for i := 0; i < len(points); i++ {
		distance := calculateDistance(points[i])
		if i < k {
			heap.Insert(Point{
				Coords: points[i],
				Dist:   distance,
			})
		} else {
			if distance < heap.Get().Dist {
				heap.Extract()
				heap.Insert(Point{
					Coords: points[i],
					Dist:   distance,
				})
			}
		}
	}

	fmt.Println(heap)

	result := make([][]int, k)
	for i := 0; i < k; i++ {
		result[i] = heap.Extract().Coords
	}

	return result
}

func main() {
	r := kClosest([][]int{
		{3, 3},
		{5, -1},
		{-2, 4},
		{1, 1},
	}, 2)

	fmt.Println(r)
}
