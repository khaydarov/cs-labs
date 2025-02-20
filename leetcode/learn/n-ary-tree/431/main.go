package main

type Node struct {
	Val      int
	Children []*Node
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	encoder func(root *Node) *TreeNode
	decoder func(root *TreeNode) *Node
}

func (c *Codec) encode(root *Node) *TreeNode {
	return c.encoder(root)
}

func (c *Codec) decode(root *TreeNode) *Node {
	return c.decoder(root)
}

func NewRecursiveCodec() *Codec {
	return &Codec{
		recursiveEncoder,
		recursiveDecoder,
	}
}

func NewIterativeCodec() *Codec {
	return &Codec{
		iterativeEncoder,
		iterativeDecoder,
	}
}
