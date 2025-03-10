package main

import (
	"math/rand"
	"time"
)

func SetupCluster(count int) Cluster {
	nodes := make([]*Node, count)
	for i := 0; i < count; i++ {
		nodes[i] = NewNode(i + 1)
	}

	for i := range nodes {
		next := (i + 1) % len(nodes)
		nodes[i].Peers = append(nodes[i].Peers, nodes[next])
		nodes[next].Peers = append(nodes[next].Peers, nodes[i])
	}

	return nodes
}

func main() {
	cluster := SetupCluster(5)

	for _, node := range cluster {
		go node.Heartbeat()
	}

	cluster.Fail(2)
	time.AfterFunc(time.Duration(rand.Intn(3)+2)*time.Second, func() {
		cluster.Recover(2)
	})

	time.AfterFunc(time.Duration(3)*time.Second, func() {
		cluster.Add(6)
	})

	cluster.Gossip()
	time.Sleep(time.Minute)
}
