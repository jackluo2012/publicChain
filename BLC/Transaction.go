package BLC

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
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
	txInput := &TXInput{[]byte{}, -1, nil, []byte{}}
	// 未消费
	txOutput := NewTXOutput(10, address)
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

func (tx *Transaction) Sign(privKey ecdsa.PrivateKey, prevTXS map[string]Transaction) {

	if tx.IsCoinbaseTransactions() {
		return
	}

	//没有找到对应的 transactions
	for _, vin := range tx.Vins {
		if prevTXS[hex.EncodeToString(vin.TxHash)].TxHash == nil {
			log.Panic("ERROR : transaction is not correct")
		}
	}

	//备份一分 tx
	txCopy := tx.TrimmedCopy()

	for inID, vin := range txCopy.Vins {
		prevTx := prevTXS[hex.EncodeToString(vin.TxHash)]
		txCopy.Vins[inID].Signature = nil
		txCopy.Vins[inID].PublicKey = prevTx.Vouts[vin.Vout].Ripemd160Hash
		txCopy.TxHash = txCopy.Hash()
		txCopy.Vins[inID].PublicKey = nil

		//签名代码
		r, s, err := ecdsa.Sign(rand.Reader, &privKey, txCopy.TxHash)
		if err != nil {
			log.Panic(err)
		}
		signature := append(r.Bytes(), s.Bytes()...)
		tx.Vins[inID].Signature = signature
	}
}

func (tx *Transaction) TrimmedCopy() Transaction {
	var inputs []*TXInput
	var outputs []*TXOutput

	for _, vin := range tx.Vins {
		inputs = append(inputs, &TXInput{vin.TxHash, vin.Vout, nil, nil})
	}

	for _, vout := range tx.Vouts {
		outputs = append(outputs, &TXOutput{vout.Value, vout.Ripemd160Hash})
	}
	txCopy := Transaction{tx.TxHash, inputs, outputs}

	return txCopy
}

func (tx *Transaction) Hash() []byte {

	txCopy := tx
	txCopy.TxHash = []byte{}

	hash := sha256.Sum256(txCopy.Serialize())

	return hash[:]
}

func (tx *Transaction) Serialize() []byte {
	var encoded bytes.Buffer

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	return encoded.Bytes()
}

// 2.转账时产生的Transaction
func NewSimpleTransaction(from, to string, amount int64, blockchain *Blockchain, txs []*Transaction) *Transaction {

	// 1.有一个函数，返回from 这个人所有的未花费交易输出所对应的Transaction

	//unUTXOs := blockchain.UnUTXOs(from)
	//unSpentTx := UnSpentTransationsWithAdress(from)

	wallets, _ := NewWallets()
	wallet := wallets.Walets[from]

	// 2.通过一个函数，返回
	money, spendableUTXODic := blockchain.FindSpendableUTXOS(from, amount, txs)

	var txInputs []*TXInput
	var txOutputs []*TXOutput

	for txHash, indexArray := range spendableUTXODic {
		txHashBytes, _ := hex.DecodeString(txHash)
		for _, index := range indexArray {
			//未花费的交易输出
			txInput := &TXInput{txHashBytes, index, nil, wallet.PublicKey}
			txInputs = append(txInputs, txInput)
		}
	}

	// 转账
	txOutput := NewTXOutput(int64(amount), to)
	txOutputs = append(txOutputs, txOutput)

	//找零
	txOutput = NewTXOutput(money-amount, from)
	//txOutput = &TXOutput{money - amount, from}
	txOutputs = append(txOutputs, txOutput)

	tx := &Transaction{[]byte(""), txInputs, txOutputs}

	//设置hash 值
	tx.HashTransaction()
	//进行签名
	blockchain.SignTransaction(tx, wallet.PrivateKey)

	return tx
}
/*
func NewUTXOTransaction(from, to string, amount int, bc *Blockchain) *Transaction {
	var inputs []TXInput
	var outputs []TXOutput

	wallets, err := NewWallets()
	if err != nil {
		log.Panic(err)
	}
	wallet := wallets.
}
*/