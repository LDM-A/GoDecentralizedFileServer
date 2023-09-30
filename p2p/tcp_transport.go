package p2p

import (
	"errors"
	"fmt"
	"log"
	"net"
	"sync"
)

// TCP Peer represents the remote node over TCP established connection.
type TCPPeer struct {
	// conn is the underlying conenction of the peer
	conn net.Conn

	// If the peer dial and retrieve a connection => outbound == true
	// If the peer accepts and retreives a connection => outbound == false
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

// Close implements the Peer interface
func (p *TCPPeer) Close() error {
	return p.conn.Close()
}

type TCPTransportOpts struct {
	ListenAddr    string
	HandshakeFunc HandshakeFunc
	Decoder       Decoder
	OnPeer        func(Peer) error
}

type TCPTransport struct {
	TCPTransportOpts
	listener net.Listener
	rpcCh    chan RPC

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
		rpcCh:            make(chan RPC),
	}

}

// Consume implements the Transport interface, which will return a read only channel.
// for reading the incomming messages received from an other Peer in the network.
func (t *TCPTransport) Consume() <-chan RPC {
	return t.rpcCh
}

// Close implements the Transport interface
func (t *TCPTransport) Close() error {
	return t.listener.Close()
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.ListenAddr)
	if err != nil {
		return err
	}

	go t.startAcceptLoop()
	log.Printf("TCP transporting listening on port %s\n", t.ListenAddr)
	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if errors.Is(err, net.ErrClosed) {
			return
		}
		if err != nil {
			fmt.Printf("TCP: accept error: %s\n", err)
		}

		fmt.Printf("new incomming connection %+v\n", conn)

		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)
	var err error
	defer func() {
		fmt.Printf("dropping peer connection: %s", err)
		conn.Close()
	}()
	if err = t.HandshakeFunc(peer); err != nil {
		return
	}

	if t.OnPeer != nil {
		if err = t.OnPeer(peer); err != nil {
			return
		}
	}

	// Read loop
	rpc := RPC{}
	for {
		err = t.Decoder.Decode(conn, &rpc)
		if err != nil {
			return
		}

		rpc.From = conn.RemoteAddr()
		t.rpcCh <- rpc
	}
}
