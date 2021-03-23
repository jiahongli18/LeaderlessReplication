package sender

import (
	"LeaderlessReplication/message"
	"encoding/gob"
	"fmt"
	"net"
)

// Dial calls the provided network address to create a TCP connection
// func Dial(i int, ID int, IDs[]string, IPs []string, ports []string, nodes map[string]net.Conn) {
func Dial(i int, ID int, nodes map[string]net.Conn) {
	ports := []string{"3000", "3001", "3002"}
	IPs := []string{"127.0.0.1", "127.0.0.1", "127.0.0.1"}
	IDs := []string{"0", "1", "2"}

	c, err := net.Dial("tcp", IPs[i]+":"+ports[i])
	if err != nil {
		fmt.Println(err)
		return
	}

	enc := gob.NewEncoder(c)
	enc.Encode(ID)
	nodes[IDs[i]] = c

	// fmt.Println(nodes)
}

// UnicastSend sends a message to other process via TCP channel
func UnicastSend(destination net.Conn, message message.Message) {
	enc := gob.NewEncoder(destination)
	enc.Encode(message)
}

// SendExit sends exit signals to connected nodes
func SendExit(nodes map[string]net.Conn, r int) {
	for _, value := range nodes {
		m := message.Message{0, r}
		UnicastSend(value, m)
	}
}
