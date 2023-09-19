package p2p

// Handshake func is something (not sure on documentation and implementation)
type HandshakeFunc func(any) error

func NOPHandshakeFunc(any) error {
	return nil
}
