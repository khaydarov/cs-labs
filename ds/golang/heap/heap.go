package heap

func Construct(capacity int) Heap {
	return Heap{
		size:     0,
		capacity: capacity,
		data:     make([]int, capacity),
	}
}

type Heap struct {
	size     int
	capacity int
	data     []int
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
func (h *Heap) Insert(k int) {
	if h.size == h.capacity {
		return
	}

	h.size++
	i := h.size - 1
	h.data[i] = k

	for i != 0 && h.data[h.Parent(i)] > h.data[i] {
		h.data[h.Parent(i)], h.data[i] = h.data[i], h.data[h.Parent(i)]
		i = h.Parent(i)
	}
}

// Get Returns first heap element: min or max value
func (h *Heap) Get() int {
	if h.size <= 0 {
		return -1
	}

	return h.data[0]
}

// Heapify adjusts heap data
func (h *Heap) Heapify(i int) {
	left := h.Left(i)
	right := h.Right(i)

	smallest := i
	if left < h.size && h.data[left] < h.data[smallest] {
		smallest = left
	}

	if right < h.size && h.data[right] < h.data[smallest] {
		smallest = right
	}

	if smallest != i {
		h.data[i], h.data[smallest] = h.data[smallest], h.data[i]
		h.Heapify(smallest)
	}
}

func (h *Heap) Extract() int {
	if h.size <= 0 {
		return -1
	}

	returnValue := h.data[0]
	h.data[0] = h.data[h.size-1]
	h.data[h.size-1] = 0
	h.size--

	h.Heapify(0)

	return returnValue
}
