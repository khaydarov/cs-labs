package main

type MapSum struct {
	trie   *Trie
	values map[string]int
}

func Constructor() MapSum {
	return MapSum{
		trie: &Trie{
			root: &Node{
				Children: make(map[rune]*Node),
			},
		},
		values: make(map[string]int),
	}
}

func (this *MapSum) Insert(key string, val int) {
	this.trie.Insert(key)
	this.values[key] = val
}

func (this *MapSum) Sum(prefix string) int {
	words := this.trie.SearchWords(prefix)
	sum := 0
	for _, word := range words {
		sum += this.values[word]
	}

	return sum
}
