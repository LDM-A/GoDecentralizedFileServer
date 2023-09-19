package main

import (
	"fmt"
	"log"

	"github.com/LDM-A/GoDecentralizedFileServer/p2p"
)

func main() {
	fmt.Println("Hello Fileserver")
	tr := p2p.NewTCPTransport(":3000")

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
}
