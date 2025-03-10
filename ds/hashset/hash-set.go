package hashset

// HashSet structure
type HashSet struct {
	size    int
	buckets [][]int
}

func Construct(size int) *HashSet {
	return &HashSet{
		size:    size,
		buckets: make([][]int, size),
	}
}

// Add appends the element to the bucket if it is not exist
func (hs *HashSet) Add(key int) {
	if hs.Contains(key) {
		return
	}

	bucket := hs.getBucket(key)
	hs.buckets[bucket] = append(hs.buckets[bucket], key)
}

// Contains returns true if key exists
func (hs *HashSet) Contains(key int) bool {
	bucket := hs.getBucket(key)
	for _, v := range hs.buckets[bucket] {
		if v == key {
			return true
		}
	}

	return false
}

// Remove removes from bucket the key
func (hs *HashSet) Remove(key int) {
	bucket := hs.getBucket(key)
	for i, v := range hs.buckets[bucket] {
		if v == key {
			hs.buckets[bucket][i] = -1
			break
		}
	}
}

// getBucket calculates hash of the key and returns index of bucket
func (hs *HashSet) getBucket(key int) int {
	return key % hs.size
}
