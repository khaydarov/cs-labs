package hashmap

type Value struct {
	Key, Value int
}

// HashMap structure
type HashMap struct {
	size    int
	buckets [][]*Value
}

func Construct(size int) *HashMap {
	return &HashMap{
		size,
		make([][]*Value, size),
	}
}

// Put — Add appends the element to the bucket if it is not exist
func (hm *HashMap) Put(key, value int) {
	bucket := hm.getBucket(key)

	updated := false
	for _, v := range hm.buckets[bucket] {
		if v.Key == key {
			v.Value = value
			updated = true
			break
		}
	}

	if !updated {
		hm.buckets[bucket] = append(hm.buckets[bucket], &Value{key, value})
	}
}

// Get returns value if key exists
func (hm *HashMap) Get(key int) int {
	bucket := hm.getBucket(key)
	for _, v := range hm.buckets[bucket] {
		if v.Key == key {
			return v.Value
		}
	}

	return -1
}

// Remove removes from bucket the key
func (hm *HashMap) Remove(key int) {
	bucket := hm.getBucket(key)
	for _, v := range hm.buckets[bucket] {
		if v.Key == key {
			v.Value = -1
			break
		}
	}
}

// getBucket calculates hash of the key and returns index of bucket
func (hm *HashMap) getBucket(key int) int {
	return key % hm.size
}
