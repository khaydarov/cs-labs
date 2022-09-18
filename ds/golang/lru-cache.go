package main

import "fmt"

// New creates new LRUCache instance
func New(capacity int) LRUCache {
	head := &item{}
	tail := &item{}

	head.next = tail
	tail.prev = head

	return LRUCache{
		storage:  make(map[int]*item, capacity),
		capacity: capacity,
		head:     head,
		tail:     tail,
	}
}

// item is doubly-linked list structure
type item struct {
	key   int
	value int

	next *item
	prev *item
}

// LRUCache is a structure that implements LRU caching strategy
type LRUCache struct {
	// key-value storage
	storage map[int]*item

	// References to doubly-linked list items: head nad tail
	head *item
	tail *item

	// Cache capacity and its current size
	capacity int
	size     int
}

// Put adds new element according to the LRU strategy
func (this *LRUCache) Put(key, value int) {
	if item, ok := this.storage[key]; ok {
		item.value = value
		this.bringToMostUsed(key)
	} else if this.capacity == 0 {
		return
	} else {
		if this.isFull() {
			this.removeLeastUsed()
		}
		this.append(key, value)
	}
}

// Get returns value from cache by key
func (this *LRUCache) Get(key int) int {
	if _, ok := this.storage[key]; !ok {
		return -1
	}

	returnValue := this.storage[key]
	this.bringToMostUsed(key)

	return returnValue.value
}

// bringToMostUsed makes key most used
func (this *LRUCache) bringToMostUsed(key int) {
	node := this.storage[key]

	mostUsed := this.tail.prev
	if node == mostUsed {
		return
	}

	node.prev.next = node.next
	node.next.prev = node.prev

	node.prev = mostUsed
	mostUsed.next = node

	this.tail.prev = node
	node.next = this.tail
}

// removeLeastUsed removes least used key
func (this *LRUCache) removeLeastUsed() {
	delete(this.storage, this.head.next.key)

	leastUsed := this.head.next
	leastUsed.next.prev = this.head
	this.head.next = leastUsed.next

	this.size--
}

// isFull returns bool if length of storage equals to capacity
func (this *LRUCache) isFull() bool {
	return len(this.storage) >= this.capacity
}

// append returns adds to the tail
func (this *LRUCache) append(key, value int) {
	node := &item{key: key, value: value}

	mostUsed := this.tail.prev
	mostUsed.next = node
	node.prev = mostUsed

	this.tail.prev = node
	node.next = this.tail

	this.storage[key] = node
	this.size++
}

func (this *LRUCache) debug() {
	curr := this.head
	for curr != nil {
		fmt.Println(
			fmt.Sprintf("curr.key: %d, curr.val: %d, next: %v, prev: %v", curr.key, curr.value, curr.next, curr.prev),
		)
		curr = curr.next
	}
}

func main() {
	// Test case
	//lru := New(2)
	//lru.Put(1,1)
	//lru.Put(2,2)
	//fmt.Println(lru.Get(1))
	//lru.Put(3,3)
	//fmt.Println(lru.Get(2))
	//lru.Put(4, 4)
	//fmt.Println(lru.Get(1))
	//fmt.Println(lru.Get(3))
	//fmt.Println(lru.Get(4))

	// Test case
	//lru := New(2)
	//fmt.Println(lru.Get(2))
	//lru.Put(2, 6)
	//fmt.Println(lru.Get(1))
	//lru.Put(1, 5)
	//lru.Put(1, 2)
	//fmt.Println(lru.Get(1))
	//fmt.Println(lru.Get(2))

	// Test case
	//lru := New(1)
	//lru.Put(2, 1)
	//fmt.Println(lru.Get(2))
	//lru.debug()
	//lru.Put(3, 2)
	//fmt.Println(lru.Get(2))
	//fmt.Println(lru.Get(3))

	//["LRUCache","put","put","put","put","put","get","put","get","get","put","get","put","put","put","get","put","get","get","get","get","put","put","get","get","get","put","put","get","put","get","put","get","get","get","put","put","put","get","put","get","get","put","put","get","put","put","put","put","get","put","put","get","put","put","get","put","put","put","put","put","get","put","put","get","put","get","get","get","put","get","get","put","put","put","put","get","put","put","put","put","get","get","get","put","put","put","get","put","put","put","get","put","put","put","get","get","get","put","put","put","put","get","put","put","put","put","put","put","put"]
	//[[4,30],[9,3],[9],[10],[10],[6,14],[3,1],[3],[10,11],[8],[2,14],[1],[5],[4],[11,4],[12,24],[5,18],[13],[7,23],[8],[12],[3,27],[2,12],[5],[2,9],[13,4],[8,18],[1,7],[6],[9,29],[8,21],[5],[6,30],[1,12],[10],[4,15],[7,22],[11,26],[8,17],[9,29],[5],[3,4],[11,30],[12],[4,29],[3],[9],[6],[3,4],[1],[10],[3,29],[10,28],[1,20],[11,13],[3],[3,12],[3,8],[10,9],[3,26],[8],[7],[5],[13,17],[2,27],[11,15],[12],[9,19],[2,15],[3,16],[1],[12,17],[9,1],[6,19],[4],[5],[5],[8,1],[11,7],[5,2],[9,28],[1],[2,2],[7,4],[4,22],[7,24],[9,26],[13,28],[11,26]]
	lru := New(10)
	lru.Put(10, 13)
	lru.Put(3, 17)
	lru.Put(6, 11)
	lru.Put(10, 5)
	lru.Put(9, 10)
	fmt.Println(lru.Get(13))
	lru.Put(2, 19)
	fmt.Println(lru.Get(2))
	fmt.Println(lru.Get(3))
	lru.Put(5, 25)
	fmt.Println(lru.Get(8))
	lru.Put(9, 22)
	lru.Put(5, 5)
	lru.Put(1, 30)
	fmt.Println(lru.Get(11))
	lru.Put(9, 12)
	fmt.Println(lru.Get(7))
	fmt.Println(lru.Get(5))
	fmt.Println(lru.Get(8))
	fmt.Println(lru.Get(9))
	lru.Put(12, 24)
	fmt.Println(lru.size)

	//lru := New(1)
	//fmt.Println(lru.Get(6))
	//fmt.Println(lru.Get(8))
	//lru.Put(12, 1)
	//fmt.Println(lru.Get(2))
	//lru.Put(15, 11)
	//lru.Put(5, 2)
	//lru.Put(1, 15)
	//lru.Put(4, 2)
	////lru.debug()
	//fmt.Println(lru.Get(5))
	//lru.Put(15, 15)
}
