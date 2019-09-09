package BLC

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/big"
	"time"
)

//数据库
const (
	dbName = "blockchain.db"
	//表名
	blockTableName = "blocks"
	//区头
	blockNewHeader = "l"
)

type Blockchain struct {
	Tip []byte //最新的区块的Hash
	DB  *bolt.DB
}

/**
 * 遍历输出所有的信息
 */
func (blc *Blockchain) PrintChain() {
	var block *Block

	var currentHash []byte = blc.Tip
	for {
		//1.取出最新的区块
		err := blc.DB.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(blockTableName))
			if b != nil {
				//获取 当前区块的字节数组
				blockBytes := b.Get(currentHash)
				//反序列化
				block = DeserializeBlock(blockBytes)

				fmt.Printf("Height: %d\n", block.Height)
				fmt.Printf("PrevBlockHash: %x\n", block.Height)
				fmt.Printf("Data: %s\n", block.Data)
				fmt.Printf("Timestamp: %s\n", time.Unix(block.TimeStamp, 0).Format("2006-01-02 03:04:05"))
				fmt.Printf("Hash: %x\n", block.Hash)
				fmt.Printf("Nonce: %d\n", block.Nonce)
			}

			return nil
		})
		if err != nil {
			log.Panic(err)
		}
		var hashInt big.Int
		hashInt.SetBytes(block.PrevBlockHash)
		if big.NewInt(0).Cmp(&hashInt) == 0 {
			break
		}

		currentHash = block.PrevBlockHash
	}
}

//增加区块到区块链里面
func (blc *Blockchain) AddBlockToBlockchain(data string) {
	//创建新区块

	// 往链里面添加区块
	err := blc.DB.Update(func(tx *bolt.Tx) error {
		//1. 获取表
		b := tx.Bucket([]byte(blockTableName))

		//2.创建新区块
		if b != nil {
			//// 先获取 最新区块

			blockBytes := b.Get(blc.Tip)
			//反序列化
			block := DeserializeBlock(blockBytes)

			newBlock := NewBlock(data, block.Height, block.Hash)
			//3.将区块序列化并且存储到数据为中
			err := b.Put(newBlock.Hash, newBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}
			//4.更新数据库里面"l"对应的hash值
			b.Put([]byte(blockNewHeader), newBlock.Hash)
			//5.更新blockchain的Tip
			blc.Tip = newBlock.Hash
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	//blc.Blocks = append(blc.Blocks, newBlock)
}

//1.创建带有创世区块的区块链

func CreateBlockchainWithGenesisBlock() *Blockchain {

	//创建或打开数据库
	db, err := bolt.Open(dbName, 0666, nil)
	if err != nil {
		log.Fatal(err)
	}
	var blockHash []byte

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b == nil {
			//创建数据库
			b, err = tx.CreateBucket([]byte(blockTableName))
			if err != nil {
				log.Panic(err)
			}
		}

		if b != nil {
			//创建创世区块
			genesisBlock := CreateGenesisBlock("Genesis Data.....")
			//将创世区块存储到表中
			err = b.Put(genesisBlock.Hash, genesisBlock.Serialize())

			if err != nil {
				log.Panic(err)
			}

			//存储最新的区块的hash
			err = b.Put([]byte(blockNewHeader), genesisBlock.Hash)
			if err != nil {
				log.Panic(err)
			}
			blockHash = genesisBlock.Hash
		}

		return nil
	})
	//返回区块链对象
	return &Blockchain{blockHash, db}

}
