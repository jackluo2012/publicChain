package BLC

import (
	"github.com/boltdb/bolt"
	"log"
)

// 迭代器
type BlockchainIterator struct {
	CurrentHash []byte //存储当前区块的Hash
	DB          *bolt.DB
}

func (bci *BlockchainIterator) Next() *Block {
	var block *Block
	err := bci.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			currentBlockBytes := b.Get(bci.CurrentHash)
			// 获取 当前迭代器里面的currentHash所对应的区块
			block = DeserializeBlock(currentBlockBytes)
			//更新迭代器里面的CurrentHash
			bci.CurrentHash = block.PrevBlockHash
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	return block
}
