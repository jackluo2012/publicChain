package main

import (
	"fmt"
	"publicChain/BLC"
)

func main() {
	//创世区块
	//blockchain := BLC.CreateBlockchainWithGenesisBlock()
	//
	////新区块
	//blockchain.AddBlockToBlockchain("Send 100RMB To jack", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	//blockchain.AddBlockToBlockchain("Send 100RMB To tom", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	//blockchain.AddBlockToBlockchain("Send 100RMB To jimes", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	//blockchain.AddBlockToBlockchain("Send 100RMB To bobi", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	//blockchain.AddBlockToBlockchain("Send 100RMB To xx", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)

	block := BLC.NewBlock("Test", 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})

	fmt.Printf("%d\n", block.Nonce)
	fmt.Printf("%x\n", block.Hash)

	proofOfWork := BLC.NewProofOfWork(block)

	fmt.Printf("%v", proofOfWork.IsValid())

	bytess := block.Serialize()
	fmt.Println(bytess)
	block = BLC.DeserializeBlock(bytess)
	fmt.Printf("%d\n", block.Nonce)
}
