package main

import "container/heap"

type Sentence struct {
	sentence string
	times    int
}

type MaxHeap []Sentence

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h MaxHeap) Less(i, j int) bool {
	if h[i].times == h[j].times {
		return h[i].sentence < h[j].sentence
	}

	return h[i].times > h[j].times
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Sentence))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]

	return x
}

type Node struct {
	Val      rune
	Children map[rune]*Node
	IsEnd    bool
}

type Trie struct {
	root *Node
}

func (t *Trie) AddSentence(sentence string) {
	current := t.root
	for _, char := range sentence {
		if _, ok := current.Children[char]; !ok {
			current.Children[char] = &Node{
				Val:      char,
				Children: make(map[rune]*Node),
			}
		}
		current = current.Children[char]
	}

	current.IsEnd = true
}

func (t *Trie) FindSentences(prefix string) []string {
	current := t.root
	for _, char := range prefix {
		if _, ok := current.Children[char]; !ok {
			return []string{}
		}

		current = current.Children[char]
	}

	var result []string
	var helper func(*Node, string)
	helper = func(node *Node, word string) {
		if node.IsEnd {
			result = append(result, word)
		}

		for _, child := range node.Children {
			helper(child, word+string(child.Val))
		}
	}

	helper(current, prefix)

	return result
}

type AutocompleteSystem struct {
	trie              *Trie
	searchString      string
	mapSentencesScore map[string]int
}

func Constructor(sentences []string, times []int) AutocompleteSystem {
	mapSentencesScore := make(map[string]int)
	for i, sentence := range sentences {
		mapSentencesScore[sentence] = times[i]
	}

	trie := &Trie{root: &Node{Children: make(map[rune]*Node)}}
	for _, sentence := range sentences {
		trie.AddSentence(sentence)
	}

	return AutocompleteSystem{
		trie,
		"",
		mapSentencesScore,
	}
}

func (this *AutocompleteSystem) Input(c byte) []string {
	if c == '#' {
		this.mapSentencesScore[this.searchString]++
		this.trie.AddSentence(this.searchString)
		this.searchString = ""
		return []string{}
	}
	this.searchString += string(c)
	sentences := this.trie.FindSentences(this.searchString)

	var searchResult []Sentence
	for _, sentence := range sentences {
		searchResult = append(searchResult, Sentence{
			sentence: sentence,
			times:    this.mapSentencesScore[sentence],
		})
	}

	maxHeap := &MaxHeap{}
	for _, sentence := range searchResult {
		*maxHeap = append(*maxHeap, sentence)
	}

	heap.Init(maxHeap)

	var result []string
	for i := 0; i < 3 && len(*maxHeap) > 0; i++ {
		result = append(result, heap.Pop(maxHeap).(Sentence).sentence)
	}

	return result
}
