package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/labstack/gommon/log"
	"publicChain/BLC"
)

func main() {

	block := BLC.NewBlock("Test", 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})

	fmt.Printf("%d\n", block.Nonce)
	fmt.Printf("%x\n", block.Hash)

	//创建或是打开数据库
	db, err := bolt.Open("my.db", 06000, nil)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("blocks"))
		if b == nil {
			b, err = tx.CreateBucket([]byte("blocks"))
			if err != nil {
				log.Panic("blocks table create faild ......")

				return err
			}
		}
		err := b.Put([]byte("l"), block.Serialize())

		if err != nil {
			return err
		}

		return nil
	})

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("blocks"))
		if b != nil {
			blockData := b.Get([]byte("l"))
			fmt.Println(blockData)
			fmt.Printf("%s\n", blockData)
			block := BLC.DeserializeBlock(blockData)
			fmt.Printf("%v\n", block)

		}

		return nil
	})

	if err != nil {
		log.Panic(err)
	}
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
