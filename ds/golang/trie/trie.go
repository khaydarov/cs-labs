package main

// Node represents a trie node
// Val: the character value of the node
// Children: a map of rune to Node
// IsEnd: a boolean to indicate if the node is the end of a word
type Node struct {
	Val      rune
	Children map[rune]*Node
	IsEnd    bool
}

type Trie struct {
	root *Node
}

// Construct creates a new Trie
// root: a node with no value and an empty map of children
func Construct() *Trie {
	return &Trie{
		root: &Node{
			Children: make(map[rune]*Node),
		},
	}
}

// Insert inserts a word into the trie
// current: a pointer to the root node
//
// Time complexity: O(n), where n is the length of the word
// Space complexity: O(n), where n is the length of the word
func (t *Trie) Insert(word string) {
	current := t.root
	for _, char := range word {
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

// StartWith checks if the trie contains a word that starts with the given prefix
// current: a pointer to the root node
//
// Time complexity: O(n), where n is the length of the prefix
// Space complexity: O(1)
func (t *Trie) StartWith(prefix string) bool {
	current := t.root
	for _, char := range prefix {
		if _, ok := current.Children[char]; !ok {
			return false
		}

		current = current.Children[char]
	}
	return true
}

// Search checks if the trie contains a word
// current: a pointer to the root node
//
// Time complexity: O(n), where n is the length of the word
// Space complexity: O(1)
func (t *Trie) Search(word string) bool {
	current := t.root
	for _, char := range word {
		if _, ok := current.Children[char]; !ok {
			return false
		}

		current = current.Children[char]
	}
	return current.IsEnd
}
