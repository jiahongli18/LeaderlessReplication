package message

import "time"

// CSMessage is a message struct that is sent between client - server
type CSMessage struct {
	key, value string
	timestamp  time.Time
}

// SSMessage is a message struct that is sent between server - server
type SSMessage struct {
	key, value string
	timestamp  time.Time
}

type ACK struct {
	ack int
}

//added this cuz it was throwing an error
type Message struct {
	State float64
	Round int
}