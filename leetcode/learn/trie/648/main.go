package main

import "strings"

type Node struct {
	Val      rune
	Children map[rune]*Node
	IsEnd    bool
}

type Trie struct {
	root *Node
}

func (t *Trie) AddWord(word string) {
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

func (t *Trie) ShortestRoot(word string) string {
	current := t.root
	for i, char := range word {
		if _, ok := current.Children[char]; !ok {
			return word
		}

		current = current.Children[char]
		if current.IsEnd {
			return word[:i+1]
		}
	}

	return word
}

func replaceWords(dictionary []string, sentence string) string {
	trie := &Trie{
		root: &Node{
			Children: make(map[rune]*Node),
		},
	}

	for _, word := range dictionary {
		trie.AddWord(word)
	}

	var result []string
	for _, word := range strings.Split(sentence, " ") {
		result = append(result, trie.ShortestRoot(word))
	}

	return strings.Join(result, " ")
}
