package main

import (
	"fmt"
	"log"
	"time"

	"github.com/LDM-A/GoDecentralizedFileServer/p2p"
)

func OnPeer(peer p2p.Peer) error {
	fmt.Println("doing some logic with the peer, outside of tcp transport")
	return nil
}
func main() {
	fmt.Println("Hello Fileserver")
	tcpTransportOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
		//TODO onPeer func
	}
	tcpTransport := p2p.NewTCPTransport(tcpTransportOpts)
	FileServerOpts := FileServerOpts{

		StorageRoot:       "3000_network",
		PathTransformFunc: CASPathTransformFunc,
		Transport:         tcpTransport,
	}
	s := NewFileServer(FileServerOpts)

	go func() {

		time.Sleep(time.Second * 3)
		s.Stop()
	}()

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

	select {}
}
