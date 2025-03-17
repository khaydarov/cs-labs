package main

import (
	"sync"

	"github.com/emirpasic/gods/sets/hashset"
)

type OrSet struct {
	adds    *hashset.Set
	removes *hashset.Set

	sync.Mutex
}

func NewOrSet() *OrSet {
	return &OrSet{
		adds:    hashset.New(),
		removes: hashset.New(),
	}
}

func (o *OrSet) Add(item string) {
	o.Lock()
	defer o.Unlock()

	o.adds.Add(item)
}

func (o *OrSet) Remove(item string) {
	o.Lock()
	defer o.Unlock()
	if !o.adds.Contains(item) {
		return
	}

	o.removes.Add(item)
}

func (o *OrSet) Merge(other *OrSet) {
	o.Lock()
	defer o.Unlock()

	o.adds = o.adds.Union(other.adds)
	o.removes = o.removes.Union(other.removes)
}

func (o *OrSet) Values() []string {
	o.Lock()
	defer o.Unlock()

	values := o.adds.Values()
	var result []string
	for _, value := range values {
		if !o.removes.Contains(value) {
			result = append(result, value.(string))
		}
	}
	return result
}

func (o *OrSet) Contains(item string) bool {
	o.Lock()
	defer o.Unlock()

	return o.adds.Contains(item) && !o.removes.Contains(item)
}

func (o *OrSet) Size() int {
	o.Lock()
	defer o.Unlock()

	return o.adds.Size() - o.removes.Size()
}
