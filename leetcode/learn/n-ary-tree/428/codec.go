package main

type Node struct {
	Val      int
	Children []*Node
}

type Codec struct {
	serializer   Serializer
	deserializer Deserializer
}

type Serializer func(root *Node) string
type Deserializer func(data string) *Node

func Constructor(serializer Serializer, deserializer Deserializer) *Codec {
	return &Codec{
		serializer,
		deserializer,
	}
}

func (this *Codec) serialize(root *Node) string {
	return this.serializer(root)
}

func (this *Codec) deserialize(data string) *Node {
	return this.deserializer(data)
}
