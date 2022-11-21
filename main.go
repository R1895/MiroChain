package main

import (
	"MicroChain/auction"
	"MicroChain/blockchain"
	"fmt"
)

func main() {
	auction.DoubleAuction()
	nt := blockchain.NodeType{}
	a := blockchain.Node{
		NT:nt,
	}
	fmt.Println("Hello world!",a)
}
