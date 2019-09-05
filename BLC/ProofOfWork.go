package BLC

import "math/big"

type ProofOfWork struct {
	Block  *Block   //当前要验证的区块
	target *big.Int // 大数存储
}

// 难度 - 挖矿 的 难度
// 256 位
// 0000 0000 0000 0000 1001 0001 0000 .... 0001

//256位Hash里面前面至少要有16个零
const targetBit = 16

func (proofOfWork *ProofOfWork) Run() ([]byte, int64) {
	return []byte{}, 0
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

	return &ProofOfWork{block, 5}
}
