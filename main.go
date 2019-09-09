package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/labstack/gommon/log"
	"publicChain/BLC"
)

func main() {

	blockchain := BLC.CreateBlockchainWithGenesisBlock()
	defer blockchain.DB.Close()

	blockchain.AddBlockToBlockchain("first blocks")
	blockchain.AddBlockToBlockchain("second blocks")
	blockchain.AddBlockToBlockchain("three  blocks")
	blockchain.PrintChain()
}

func boltTest() {
	db, err := bolt.Open("my.db", 06000, nil)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	//创建表 .
	err = db.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte("BlockBucket"))
		if b == nil {
			//创建BlockBucket表
			b, err = tx.CreateBucket([]byte("BlockBucket"))
			if err != nil {
				return fmt.Errorf("Create bucket: %s", err)
			}
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
