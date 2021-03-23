package receiver

import (
	"LeaderlessReplication/message"
	"encoding/gob"
	"fmt"
	"net"
	"os"
)

func Listen(port string, ID int, IDs []string, nodes map[string]net.Conn) {
	for {
		l, err := net.Listen("tcp", ":"+port)
		if err != nil {
			fmt.Println(err)
			return
		}

		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		var temp string
		dec := gob.NewDecoder(c)
		dec.Decode(&temp)
		nodes[temp] = c
		l.Close()
	}
}

// ServerListen accepts incoming dials and establishes connections with all other nodes. Used by server.go
func ServerListen(port string, nodes map[string]net.Conn) {
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println(err)
		return
	}

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	var ID string
	dec := gob.NewDecoder(c) // read from the channel
	dec.Decode(&ID)
	nodes[ID] = c

	// fmt.Println(ID, c)
	l.Close()
}

// ListenForExit waits for the command from server letting nodes know consensus has been reached and to exit
func ListenForExit(nodes map[string]net.Conn) {
	var kill message.Message
	conn := nodes["0"]
	UnicastReceive(conn, &kill)
	if kill.State == 0 {
		fmt.Printf("Consensus reached in round %v. Exiting...\n", kill.Round)
		os.Exit(0)
	}
}

// UnicastReceive receives a message from the TCP channel
func UnicastReceive(source net.Conn, message *message.Message) {
	dec := gob.NewDecoder(source)
	dec.Decode(message)
}
