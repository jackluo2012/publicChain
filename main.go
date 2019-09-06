package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/labstack/gommon/log"
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

	//
	boltTest()
}

func boltTest() {
	db, err := bolt.Open("my.db", 06000, nil)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	//创建表 .
	err = db.Update(func(tx *bolt.Tx) error {

		//创建BlockBucket表
		b, err := tx.CreateBucket([]byte("BlockBucket"))
		if err != nil {
			return fmt.Errorf("Create bucket: %s", err)
		}
		//往表里面存储数据
		if b != nil {
			b.Put([]byte("l"), []byte("Send 100 BTC To 强哥......"))
			if err != nil {
				log.Panic("数据存储失败")
			}
		}

		return nil
	})

	// 更新失败
	if err != nil {
		log.Panic(err)
	}
}
