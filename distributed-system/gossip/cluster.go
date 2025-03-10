package main

import (
	"fmt"
	"time"
)

type Cluster []*Node

func (n *Cluster) Fail(i int) {
	(*n)[i].Fail()
}

func (n *Cluster) Recover(i int) {
	(*n)[i].Recover()
}

func (n *Cluster) Add(i int) {
	newNode := NewNode(i)

	l := len(*n)
	target := (*n)[l-1]
	target.Peers = append(target.Peers, newNode)
	newNode.Peers = append(newNode.Peers, target)

	go newNode.Heartbeat()
	fmt.Printf("Added Node %d to the network\n", i)
}

func (n *Cluster) Gossip() {
	for _, node := range *n {
		node.Gossip()
	}

	time.Sleep(time.Second)
}
