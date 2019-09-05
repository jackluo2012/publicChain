package main

import (
	"fmt"
	"publicChain/BLC"
)

func main() {
	genesisBlockchain := BLC.CreateBlockchainWithGenesisBlock()

	fmt.Print(genesisBlockchain)
	fmt.Print(genesisBlockchain.Blocks)
	fmt.Print(genesisBlockchain.Blocks[0].Data)
}
