package BLC

import "bytes"

type TXInput struct {
	//
	TxHash []byte
	// 2.存储的TXOutput 里面的索引
	Vout int
	//用户名
	Signature []byte //数字签名

	PublicKey []byte //公钥,钱包里面
}

/**
 * 判断当前的消费是谁的钱 是否是 一一对应
 */
// 判断当前的消费是谁的钱
func (txInput *TXInput) UnLockRipemd160Hash(ripemd160Hash []byte) bool {

	publicKey := Ripemd160Hash(txInput.PublicKey)

	return bytes.Compare(publicKey, ripemd160Hash) == 0
}
