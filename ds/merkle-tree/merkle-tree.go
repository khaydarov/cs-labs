package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
)

type HashValue []byte

func (v HashValue) String() string {
	return hex.EncodeToString(v)
}

type Hasher interface {
	Hash(input []byte) HashValue
	HashPair(left, right []byte) HashValue
}

type MerkleTree struct {
	root      *Node
	leaves    []*Node
	leavesMap map[string]int

	hasher Hasher
}

func New(values []string, hasher Hasher) *MerkleTree {
	leaves := make([]*Node, 0, len(values))
	for _, v := range values {
		leaves = append(leaves, &Node{
			Hash:  hasher.Hash([]byte(v)),
			Value: v,
		})
	}

	leavesMap := make(map[string]int)
	for i, l := range leaves {
		leavesMap[l.Hash.String()] = i
	}

	return &MerkleTree{
		root:      buildRoot(leaves, hasher),
		leaves:    leaves,
		leavesMap: leavesMap,
		hasher:    hasher,
	}
}

func buildRoot(leaves []*Node, hasher Hasher) *Node {
	switch len(leaves) {
	case 0:
		return nil
	case 1:
		return leaves[0]
	case 2:
		left, right := leaves[0], leaves[1]

		parent := &Node{
			Left:  left,
			Right: right,
			Hash:  hasher.HashPair(left.Hash, right.Hash),
			Value: fmt.Sprintf("%s_%s", left.Value, right.Value),
		}

		left.Parent, right.Parent = parent, parent

		return parent
	default:
		leaves = leaves[:len(leaves):len(leaves)]

		if len(leaves)%2 == 1 {
			leaves = append(leaves, leaves[len(leaves)-1])
		}

		half := len(leaves) / 2
		left := buildRoot(leaves[:half], hasher)
		right := buildRoot(leaves[half:], hasher)

		parent := &Node{
			Left:  left,
			Right: right,
			Hash:  hasher.HashPair(left.Hash, right.Hash),
			Value: fmt.Sprintf("%s_%s", left.Value, right.Value),
		}

		left.Parent, right.Parent = parent, parent

		return parent
	}
}

func (t *MerkleTree) Verify(value string) bool {
	if t == nil || t.root == nil {
		return false
	}

	node := t.findNode(value)

	if node == nil {
		return false
	}

	h := t.hasher.Hash([]byte(value))

	for parent, current := node.Parent, node; parent != nil; parent, current = parent.Parent, parent {
		if parent.Left == current {
			h = t.hasher.HashPair(h, parent.Right.Hash)
		} else {
			h = t.hasher.HashPair(parent.Left.Hash, h)
		}
	}

	return bytes.Equal(h, t.root.Hash)
}

func (t *MerkleTree) findNode(value string) *Node {
	hash := t.hasher.Hash([]byte(value))

	if index, ok := t.leavesMap[hash.String()]; ok {
		return t.leaves[index]
	}

	return nil
}
