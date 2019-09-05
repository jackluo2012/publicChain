package main

import (
	"publicChain/BLC"
)

func main() {
	//创世区块
	blockchain := BLC.CreateBlockchainWithGenesisBlock()

	//新区块
	blockchain.AddBlockToBlockchain("Send 100RMB To jack", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	blockchain.AddBlockToBlockchain("Send 100RMB To tom", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	blockchain.AddBlockToBlockchain("Send 100RMB To jimes", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	blockchain.AddBlockToBlockchain("Send 100RMB To bobi", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	blockchain.AddBlockToBlockchain("Send 100RMB To xx", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)

}
