package BLC

import "fmt"
//  创建创世区块

func (cli *Cli) createGenesisBlockchain(address string) {

	// 1. 创建coinBase

	blockchain := CreateBlockchainWithGenesisBlock(address)
	defer blockchain.DB.Close()

	fmt.Printf("地址：%s", address)
}
