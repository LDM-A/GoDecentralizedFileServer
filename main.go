package main

import (
	"fmt"
	"log"

	"github.com/LDM-A/GoDecentralizedFileServer/p2p"
)

func main() {
	fmt.Println("Hello Fileserver")
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
	}
	tr := p2p.NewTCPTransport(tcpOpts)

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
}
