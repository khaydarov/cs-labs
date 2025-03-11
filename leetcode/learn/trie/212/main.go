package main

type TrieNode struct {
	Value    rune
	Word     string
	Children map[rune]*TrieNode
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{root: &TrieNode{Children: make(map[rune]*TrieNode)}}
}

func findWords(board [][]byte, words []string) []string {
	trie := NewTrie()
	for _, word := range words {
		current := trie.root
		for _, char := range word {
			if _, ok := current.Children[char]; !ok {
				current.Children[char] = &TrieNode{
					Value:    char,
					Children: make(map[rune]*TrieNode),
				}
			}
			current = current.Children[char]
		}

		current.Word = word
	}

	var result []string
	var backtrack func(row, col int, node *TrieNode)
	backtrack = func(row, col int, node *TrieNode) {
		letter := board[row][col]
		current := node.Children[rune(letter)]

		if current.Word != "" {
			result = append(result, current.Word)
			current.Word = ""
		}

		board[row][col] = '#'
		dirs := [][]int{
			{0, 1},
			{1, 0},
			{0, -1},
			{-1, 0},
		}

		for _, dir := range dirs {
			newRow := row + dir[0]
			newCol := col + dir[1]

			if newRow < 0 || newRow >= len(board) || newCol < 0 || newCol >= len(board[0]) {
				continue
			}

			if _, ok := current.Children[rune(board[newRow][newCol])]; ok {
				backtrack(newRow, newCol, current)
			}
		}

		board[row][col] = letter
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if _, ok := trie.root.Children[rune(board[i][j])]; ok {
				backtrack(i, j, trie.root)
			}
		}
	}

	return result
}
