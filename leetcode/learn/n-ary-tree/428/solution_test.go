package main

import (
	"fmt"
	"testing"
)

func buildTestCase1() *Node {
	one := &Node{Val: 1}
	two := &Node{Val: 2}
	three := &Node{Val: 3}
	four := &Node{Val: 4}
	five := &Node{Val: 5}

	one.Children = []*Node{two, three}
	two.Children = []*Node{four}
	four.Children = []*Node{five}

	return one
}

var testCases = []struct {
	root *Node
}{
	{
		root: buildTestCase1(),
	},
	{
		root: nil,
	},
}

func TestSolution1(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			codec := NewBinaryCodec()
			serialized := codec.serialize(tc.root)
			deserialized := codec.deserialize(serialized)

			check := codec.serialize(deserialized)
			if check != serialized {
				t.Error("check failed")
			}
		})
	}
}

func TestSolution2(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			codec := NewNaryCodec()
			serialized := codec.serialize(tc.root)
			deserialized := codec.deserialize(serialized)

			check := codec.serialize(deserialized)
			if check != serialized {
				t.Error("check failed")
			}
		})
	}
}

func TestSolution3(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test/%d", i), func(t *testing.T) {
			codec := NewNaryCodec1()
			serialized := codec.serialize(tc.root)
			deserialized := codec.deserialize(serialized)

			check := codec.serialize(deserialized)
			if check != serialized {
				t.Error("check failed")
			}
		})
	}
}
