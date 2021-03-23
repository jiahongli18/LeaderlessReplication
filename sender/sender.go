package sender

// Dial calls the provided network address to create a TCP connection
func Dial(i int, ID string, IDs, IPs []string, ports []string, nodes map[string]net.Conn) {
	c, err := net.Dial("tcp", IPs[i]+":"+ports[i])
	if err != nil {
		fmt.Println(err)
		return
	}

	enc := gob.NewEncoder(c)
	enc.Encode(ID)
	nodes[IDs[i]] = c
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