package main

import (
	"flag"
	"log"
	"net/http"
)

func runServer(port string) {
	node := NewNode(NodeId(RandStringBytes(6)))
	http.HandleFunc("/", node.handleIndex)
	http.HandleFunc("/increment", node.handleIncrement)
	http.HandleFunc("/value", node.handleGetValue)
	http.HandleFunc("/merge", node.handleMerge)
	http.HandleFunc("/add-peer", node.handleAddPeer)
	http.HandleFunc("/peers", node.handleGetPeers)

	go node.gossipLoop()

	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func main() {
	var port string
	flag.StringVar(&port, "port", "8080", "Node port")
	flag.Parse()

	runServer(port)
}
