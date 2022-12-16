package main

import (
	"fmt"
)

type Set struct {
	hmap map[string]bool
}

func (s *Set) Add(v string) {
	if _, ok := s.hmap[v]; !ok {
		s.hmap[v] = true
	}
}

func (s *Set) Exists(v string) bool {
	if _, ok := s.hmap[v]; ok {
		return true
	}

	return false
}

func (s *Set) Empty() bool {
	return len(s.hmap) == 0
}

func (s *Set) Size() int {
	return len(s.hmap)
}

type ValidWordAbbr struct {
	dictionary map[string]Set
}

func Constructor(dictionary []string) ValidWordAbbr {
	hmap := make(map[string]Set)
	for _, word := range dictionary {
		abbr := abbreviate(word)

		wordSet, ok := hmap[abbr]
		if !ok {
			newSet := Set{
				make(map[string]bool),
			}
			newSet.Add(word)
			hmap[abbr] = newSet
		} else {
			wordSet.Add(word)
		}
	}

	return ValidWordAbbr{
		hmap,
	}
}

func (this *ValidWordAbbr) IsUnique(word string) bool {
	wordAbbr := abbreviate(word)

	abbrSet, ok := this.dictionary[wordAbbr]
	if !ok || abbrSet.Empty() {
		return true
	}

	if abbrSet.Exists(word) && abbrSet.Size() == 1 {
		return true
	}

	return false
}

func abbreviate(word string) string {
	if len(word) > 2 {
		return word[:1] + fmt.Sprintf("%v", len(word)-2) + word[len(word)-1:]
	}

	return word
}

func main() {
	q := Constructor([]string{"deer", "door", "cake", "card"})
	q.IsUnique("coke")

}
