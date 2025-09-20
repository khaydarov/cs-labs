package main

import (
	"encoding/hex"
	"fmt"
)

type Node struct {
	Left   *Node
	Right  *Node
	Parent *Node

	Hash   HashValue
	Value  string
	IsCopy bool
}

func (n *Node) Copy() *Node {
	return &Node{
		Left:   n.Left,
		Right:  n.Right,
		Hash:   n.Hash,
		Value:  n.Value,
		IsCopy: true,
	}
}

func (n *Node) String() string {
	return fmt.Sprintf("{%s, %s}", hex.EncodeToString(n.Hash), n.Value)
}
