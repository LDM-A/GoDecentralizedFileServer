package p2p

// Handshake func is something (not sure on documentation and implementation)
type HandshakeFunc func(Peer) error

func NOPHandshakeFunc(Peer) error {
	return nil
}
