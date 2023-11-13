package p2p

import "net"

// Peer is an interface that represents the remote node
type Peer interface {
	Send([]byte) error
	net.Conn
}

// Transport is anything that can handle the communication
// between nodes in the network.
// This can be of the form (TCP, UDP, websockets)
type Transport interface {
	Dial(string) error
	ListenAndAccept() error
	Consume() <-chan RPC
	Close() error
}
