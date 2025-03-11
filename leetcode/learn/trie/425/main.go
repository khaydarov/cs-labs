package main

type TrieNode struct {
	Val      rune
	Children map[rune]*TrieNode
}

type Trie struct {
	root *TrieNode
}

func buildTrie(words []string) *Trie {
	trie := &Trie{root: &TrieNode{Children: make(map[rune]*TrieNode)}}

	for _, word := range words {
		current := trie.root
		for _, char := range word {
			if current.Children[char] == nil {
				current.Children[char] = &TrieNode{Children: make(map[rune]*TrieNode)}
			}

			current = current.Children[char]
		}
	}

	return trie
}

func (t *Trie) WithPrefix(prefix string) []string {
	var words []string

	current := t.root
	for _, char := range prefix {
		if current.Children[char] == nil {
			return words
		}
		current = current.Children[char]
	}

	var helper func(node *TrieNode, word string)
	helper = func(node *TrieNode, word string) {
		if len(node.Children) == 0 {
			words = append(words, word)
			return
		}

		for char, child := range node.Children {
			helper(child, word+string(char))
		}
	}

	helper(current, prefix)
	return words
}

type WordSquare []string

func (w WordSquare) getPrefix(index int) string {
	var prefix string
	for _, word := range w {
		prefix += string(word[index])
	}

	return prefix
}

func (w *WordSquare) Push(word string) {
	*w = append(*w, word)
}

func (w *WordSquare) Pop() {
	*w = (*w)[:len(*w)-1]
}

// TC: O(N * 26^L * L)
// SC: O(N * L)
func wordSquares(words []string) [][]string {
	prefixTrie := buildTrie(words)

	var result [][]string
	var backtrack func(wordSquare WordSquare, index int)
	backtrack = func(wordSquare WordSquare, index int) {
		if index == len(words[0]) {
			copyWordSquare := make(WordSquare, len(wordSquare))
			copy(copyWordSquare, wordSquare)
			result = append(result, copyWordSquare)
			return
		}

		prefix := wordSquare.getPrefix(index)
		for _, candidate := range prefixTrie.WithPrefix(prefix) {
			wordSquare.Push(candidate)
			backtrack(wordSquare, index+1)
			wordSquare.Pop()
		}
	}

	for _, word := range words {
		var wordSquare WordSquare
		wordSquare.Push(word)
		backtrack(wordSquare, 1)
	}

	return result
}
