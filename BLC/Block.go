package BLC

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

//定义区块
type Block struct {
	//1.区块高度 (编号)
	Height int64
	//2.上一个 区块HASH
	PrevBlockHash []byte
	//3.交易数据
	Data []byte
	//4.时间戮
	TimeStamp int64
	//5.Hash
	Hash []byte
}

func (b *Block) SetHash() {

	// 1.height 将转换成字节数组 []byte
	heightBytes := IntToHex(b.Height)

	fmt.Println(heightBytes)
	// 2.将时间戮 []byte
	// 2 ~ 36
	timeString := strconv.FormatInt(b.TimeStamp, 2)

	timBytes := []byte(timeString)
	fmt.Println(timBytes)
	//3.拼接所有属性
	blockBytes := bytes.Join([][]byte{heightBytes, b.PrevBlockHash, b.Data, timBytes, b.Hash}, []byte{})
	//4.生成Hash //固定大小的字节数组
	hash := sha256.Sum256(blockBytes)
	//切片数字组 ，一个是固定，一个是不固定
	b.Hash = hash[:]
}

//1.创建新的区块
func NewBlock(data string, height int64, prevBlockHash []byte) *Block {
	//创建区块
	block := &Block{height, prevBlockHash, []byte(data), time.Now().Unix(), nil}

	// 设置Hash
	block.SetHash()

	return block
}
