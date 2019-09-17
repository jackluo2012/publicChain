package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
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

// 判断是否创世区块的交易
func (tx *Transaction) IsCoinbaseTransactions() bool {

	return len(tx.Vins[0].TxHash) == 0 && tx.Vins[0].Vout == -1
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

// 2.转账时产生的Transaction
func NewSimpleTransaction(from, to string, amount int64, blockchain *Blockchain) *Transaction {

	// 1.有一个函数，返回from 这个人所有的未花费交易输出所对应的Transaction

	//unUTXOs := blockchain.UnUTXOs(from)
	//unSpentTx := UnSpentTransationsWithAdress(from)

	// 2.通过一个函数，返回
	money, spendableUTXODic := blockchain.FindSpendableUTXOS(from, amount)

	var txInputs []*TXInput
	var txOutputs []*TXOutput

	for txHash, indexArray := range spendableUTXODic {
		txHashBytes, _ := hex.DecodeString(txHash)
		for _, index := range indexArray {
			//未花费的交易输出
			txInput := &TXInput{txHashBytes, index, from}
			txInputs = append(txInputs, txInput)
		}
	}

	// 转账
	txOutput := &TXOutput{amount, to}
	txOutputs = append(txOutputs, txOutput)

	//找零
	txOutput = &TXOutput{money - amount, from}
	txOutputs = append(txOutputs, txOutput)

	tx := &Transaction{[]byte(""), txInputs, txOutputs}

	//设置hash 值
	tx.HashTransaction()

	return tx
}
