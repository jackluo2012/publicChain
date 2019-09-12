package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
	"time"
)

//定义区块
type Block struct {
	//1.区块高度 (编号)
	Height int64
	//2.上一个 区块HASH
	PrevBlockHash []byte
	//3.交易数据
	Txs []*Transaction
	//4.时间戮
	TimeStamp int64
	//5.Hash
	Hash []byte
	//6.Nonce
	Nonce int64
}

// 将区块序列化成字节数组
func (block *Block) HashTransactions() []byte {

	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range block.Txs {
		txHashes = append(txHashes, tx.TxHash)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))
	return txHash[:]
}

//将区块序列化成字节数组
func (block *Block) Serialize() []byte {
	var result bytes.Buffer

	//打包
	encoder := gob.NewEncoder(&result)
	//序列化
	err := encoder.Encode(block)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

//反序化
func DeserializeBlock(blockBytes []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(blockBytes))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}

//1.创建新的区块
func NewBlock(txs []*Transaction, height int64, prevBlockHash []byte) *Block {
	//创建区块
	block := &Block{height, prevBlockHash, txs, time.Now().Unix(), nil, 0}

	// 调用工作量证明的方法并且返回有效的hash和Nonce值
	//返回有效的hash 和nonce值
	pow := NewProofOfWork(block)
	//
	// 000000
	hash, nonce := pow.Run()
	fmt.Println()
	block.Nonce = nonce
	block.Hash = hash[:]
	return block
}

// 2.单独写一个方法，生成创世区块
func CreateGenesisBlock(txs []*Transaction) *Block {

	return NewBlock(txs, 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}
