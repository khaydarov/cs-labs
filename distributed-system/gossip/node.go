package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Node struct {
	ID     int
	Peers  []*Node
	Active bool
	Delay  time.Duration

	sync.Mutex
}

func NewNode(id int) *Node {
	return &Node{
		ID:     id,
		Active: true,
		Delay:  time.Duration(3) * time.Second,
	}
}

func (n *Node) Heartbeat() {
	for {
		time.Sleep(time.Second)
		n.Mutex.Lock()
		if !n.Active {
			n.Mutex.Unlock()
			continue
		}

		fmt.Printf("Node %d is executing heartbeat\n", n.ID)
		for _, peer := range n.Peers {
			if !peer.Active {
				fmt.Printf("Node %d detected failure of Node %d\n", n.ID, peer.ID)
			}
		}
		n.Mutex.Unlock()
	}
}

func (n *Node) Recover() {
	n.Mutex.Lock()
	n.Active = true
	n.Mutex.Unlock()
	fmt.Printf("Node %d recovered!\n", n.ID)
}

func (n *Node) Fail() {
	n.Mutex.Lock()
	n.Active = false
	n.Mutex.Unlock()
	fmt.Printf("Node %d failed!\n", n.ID)
}

func (n *Node) Gossip() {
	if !n.Active {
		return
	}

	if len(n.Peers) == 0 {
		return
	}

	n.Mutex.Lock()
	defer n.Mutex.Unlock()

	sendCount := rand.Intn(len(n.Peers)/2+1) + 1
	sentNodes := make(map[int]bool)

	time.Sleep(n.Delay)
	for i := 0; i < sendCount; i++ {
		target := n.Peers[rand.Intn(len(n.Peers))]
		if !sentNodes[target.ID] {
			if target.Active {
				fmt.Printf("Node %d shared data with Node %d\n", n.ID, target.ID)
			}
			sentNodes[target.ID] = true
		}
	}
}
