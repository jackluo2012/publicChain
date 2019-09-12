package BLC

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

type ProofOfWork struct {
	Block  *Block   //当前要验证的区块
	target *big.Int // 大数存储
}

func (pow *ProofOfWork) prepareData(nonce int64) []byte {
	data := bytes.Join([][]byte{pow.Block.PrevBlockHash,
		pow.Block.HashTransactions(),
		IntToHex(pow.Block.TimeStamp),
		IntToHex(int64(targetBit)),
		IntToHex(int64(nonce)),
		IntToHex(int64(pow.Block.Height)),
	},
		[]byte{})
	return data
}

// 难度 - 挖矿 的 难度
// 256 位
// 0000 0000 0000 0000 1001 0001 0000 .... 0001

//256位Hash里面前面至少要有16个零
const targetBit = 16

/**
 * 检查 工作量算法，是否有效
 */
func (proofOfWork *ProofOfWork) IsValid() bool {

	// 1.proofOfWork.Blck.Heigh
	// 2.proofOfWork.Target
	var hashInt big.Int
	hashInt.SetBytes(proofOfWork.Block.Hash)

	if proofOfWork.target.Cmp(&hashInt) == 1 {
		return true
	}

	return false
}

func (proofOfWork *ProofOfWork) Run() ([]byte, int64) {
	// 1.将block 的属性拼接成字节数组

	// 2.生成hash

	//3.判断hash有效性，如果满足条件，跳出循环
	var nonce int64
	var hashInt big.Int //存储我们新生成的hash
	var hash [32]byte
	for {
		dataBytes := proofOfWork.prepareData(nonce)

		//生成hash
		hash = sha256.Sum256(dataBytes)
		fmt.Printf("\r%x", hash)
		//将hash 存储到hashInt
		hashInt.SetBytes(hash[:])

		//判断hashInt 是否小于Block里面的target
		if proofOfWork.target.Cmp(&hashInt) == 1 {
			break
		}

		nonce++
	}

	return hash[:], nonce
}

// 创建新的工作量工作证明对象

func NewProofOfWork(block *Block) *ProofOfWork {
	// 1.big.Int对象 1
	// 2
	//0000 0001
	// 8 - 2 = 6
	// 0100 0000  64

	// 1.创建一个如始值为1的target
	target := big.NewInt(1)
	// 2.左移256 - targetBit
	target = target.Lsh(target, 256-targetBit)

	return &ProofOfWork{block, target}
}
