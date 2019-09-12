package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

//UTXO
type Transaction struct {
	// 1. 交易 Hash
	TxHash []byte

	// 2. 输入
	Vins []*TXInput

	// 3. 输出
	Vouts []*TXOutput
}

//  1.Transaction 创建分两种情况

// 1.创世区块创建时的Transaction

func NewCoinBaseTransaction(address string) *Transaction {

	// 代表消费信息
	txInput := &TXInput{[]byte{}, -1, "Genesis Block"}
	// 未消费
	txOutput := &TXOutput{10, address}

	txCoinbase := &Transaction{[]byte{}, []*TXInput{txInput}, []*TXOutput{txOutput}}

	//设置hash值
	txCoinbase.HashTransaction()

	return txCoinbase

}

// 2.转账时产生的Transaction

/**
 * 序列化一个区块
 */
func (tx *Transaction) HashTransaction() {
	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(tx)
	if err != nil {
		log.Panic(err)
	}

	hash := sha256.Sum256(result.Bytes())

	tx.TxHash = hash[:]
}
