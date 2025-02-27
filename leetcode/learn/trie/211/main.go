package main

type Node struct {
	Val      rune
	Children map[rune]*Node
	IsEnd    bool
}

type Trie struct {
	root *Node
}

func (t *Trie) Add(word string) {
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

func (t *Trie) Search(word string) bool {
	var helper func(*Node, string) bool
	helper = func(node *Node, searchString string) bool {
		if searchString == "" {
			return node.IsEnd
		}

		v := []rune(searchString)
		if v[0] != '.' {
			if _, ok := node.Children[v[0]]; !ok {
				return false
			}

			return helper(node.Children[v[0]], string(v[1:]))
		}

		result := false
		for _, child := range node.Children {
			result = result || helper(child, string(v[1:]))
		}

		return result
	}

	return helper(t.root, word)
}

func NewTrie() Trie {
	return Trie{
		root: &Node{
			Children: make(map[rune]*Node),
		},
	}
}

type WordDictionary struct {
	trie *Trie
}

func Constructor() WordDictionary {
	return WordDictionary{
		trie: &Trie{
			root: &Node{
				Children: make(map[rune]*Node),
			},
		},
	}
}

func (this *WordDictionary) AddWord(word string) {
	this.trie.Add(word)
}

func (this *WordDictionary) Search(word string) bool {
	return this.trie.Search(word)
}
