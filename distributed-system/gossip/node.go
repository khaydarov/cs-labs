package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type NodeId string

type Node struct {
	ID      NodeId
	Counter *GCounter
	Peers   map[string]bool

	sync.Mutex
}

func NewNode(id NodeId) *Node {
	return &Node{
		ID:      id,
		Counter: NewGCounter(),
		Peers:   make(map[string]bool),
	}
}

func (n *Node) String() string {
	return string(n.ID)
}

func (n *Node) handleIndex(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hello from Node `%s`!", n.ID)))
}

func (n *Node) handleIncrement(w http.ResponseWriter, r *http.Request) {
	n.Lock()
	defer n.Unlock()

	n.Counter.Increment(n.ID)
	log.Printf("Node `%s` incremented to %d", n.ID, n.Counter.Value())

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%d", n.Counter.Value())))
}

func (n *Node) handleGetValue(w http.ResponseWriter, r *http.Request) {
	n.Lock()
	defer n.Unlock()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%d", n.Counter.Value())))
}

func (n *Node) handleMerge(w http.ResponseWriter, r *http.Request) {
	var remote GCounter
	if err := json.NewDecoder(r.Body).Decode(&remote); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	n.Lock()
	defer n.Unlock()

	n.Counter.Merge(&remote)
	log.Printf("Node `%s` merged with %v", n.ID, remote)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%d", n.Counter.Value())))
}

func (n *Node) handleAddPeer(w http.ResponseWriter, r *http.Request) {
	var peer struct {
		Peer string `json:"peer"`
	}

	if err := json.NewDecoder(r.Body).Decode(&peer); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	n.Lock()
	defer n.Unlock()

	n.Peers[peer.Peer] = true
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Added peer `%s`", peer.Peer)))
}

func (n *Node) handleGetPeers(w http.ResponseWriter, r *http.Request) {
	n.Lock()
	defer n.Unlock()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%v", n.Peers)))
}

func (n *Node) gossipLoop() {
	for {
		time.Sleep(3 * time.Second)

		n.Lock()
		peers := make([]string, 0, len(n.Peers))
		for peer := range n.Peers {
			peers = append(peers, peer)
		}
		n.Unlock()

		if len(peers) == 0 {
			log.Printf("Node `%s`: No peers to gossip to", n.ID)
			continue
		}

		target := peers[rand.Intn(len(peers))]
		go n.gossipWith(target)
	}
}

func (n *Node) gossipWith(peer string) {
	n.Lock()
	data, _ := json.Marshal(n.Counter)
	n.Unlock()

	log.Printf("Gossiping to %s", peer)
	resp, err := http.Post(
		fmt.Sprintf("%s/merge", peer),
		"application/json",
		bytes.NewBuffer(data),
	)

	if err != nil {
		log.Printf("Error gossiping to %s: %v", peer, err)
		return
	}
	defer resp.Body.Close()
}
