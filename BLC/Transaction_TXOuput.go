package BLC

import "bytes"

type TXOutput struct {
	//多少钱
	Value int64

	Ripemd160Hash []byte //由公钥转换而成的
}

/**
 * 判断当前的消费是谁的钱 是否是 一一对应
 * 进行解锁
 */
func (txOutput *TXOutput) UnLockWithAdress(address string) bool {

	publicKeyHash := Base58Decode([]byte(address))

	hash160 := publicKeyHash[1 : len(publicKeyHash)-4]

	return bytes.Compare(txOutput.Ripemd160Hash, hash160) == 0
}

/**
 * 上锁的目的是为了设置Hash160的值
 */

func (txOutput *TXOutput) Lock(address string) {

	publicKeyHash := Base58Decode([]byte(address))

	txOutput.Ripemd160Hash = publicKeyHash[1 : len(publicKeyHash)-4]

}

func NewTXOutput(value int64, address string) *TXOutput {

	txOutput := &TXOutput{value, nil}

	// 设置Ripemd160Hash
	txOutput.Lock(address)
	return txOutput
}
