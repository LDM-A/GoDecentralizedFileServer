package main

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/LDM-A/GoDecentralizedFileServer/p2p"
)

func makeServer(listenAddr string, nodes ...string) *FileServer {
	tcpTransportOpts := p2p.TCPTransportOpts{
		ListenAddr:    listenAddr,
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},

		//TODO onPeer func
	}
	tcpTransport := p2p.NewTCPTransport(tcpTransportOpts)
	FileServerOpts := FileServerOpts{

		StorageRoot:       listenAddr + "_network",
		PathTransformFunc: CASPathTransformFunc,
		Transport:         tcpTransport,
		BootstrapNodes:    nodes,
	}
	s := NewFileServer(FileServerOpts)
	tcpTransport.OnPeer = s.OnPeer
	return s

}
func OnPeer(peer p2p.Peer) error {
	fmt.Println("doing some logic with the peer, outside of tcp transport")
	return nil
}
func main() {
	s1 := makeServer(":3000", "")
	s2 := makeServer(":4000", ":3000")
	go func() {
		if err := s1.Start(); err != nil {
			log.Fatal(err)
		}
	}()
	time.Sleep(2 * time.Second)
	go s2.Start()
	time.Sleep(2 * time.Second)
	data := bytes.NewReader([]byte("big file here"))
	s2.StoreData("key", data)
	select {}
}
