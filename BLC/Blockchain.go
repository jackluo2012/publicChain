package BLC

import (
	"encoding/hex"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/big"
	"os"
	"strconv"
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

//创建区块链的迭代器
func (blockchain *Blockchain) Iterator() *BlockchainIterator {
	return &BlockchainIterator{blockchain.Tip, blockchain.DB}
}

//判断数据库是否存在
func dbExists() bool {
	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		return false
	}
	return true
}

/**
 * 遍历输出所有的信息
 */
func (blc *Blockchain) PrintChain() {
	blockchainIterator := blc.Iterator()
	for {
		block := blockchainIterator.Next()

		fmt.Printf("Height: %d\n", block.Height)
		fmt.Printf("PrevBlockHash: %x\n", block.Height)

		fmt.Printf("Timestamp: %s\n", time.Unix(block.TimeStamp, 0).Format("2006-01-02 03:04:05"))
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Nonce: %d\n", block.Nonce)

		for _, tx := range block.Txs {
			fmt.Printf("%s\n", tx.TxHash)
			fmt.Println("Vins:")
			for _, in := range tx.Vins {
				fmt.Printf("%s\n", in.TxHash)
				fmt.Printf("%d\n", in.Vout)
			}
			fmt.Println("Vouts:")
			for _, out := range tx.Vouts {
				fmt.Println(out.Value)
				fmt.Println(out.Ripemd160Hash)
			}
		}

		var hashInt big.Int
		//把字节数组转换成数字
		hashInt.SetBytes(block.PrevBlockHash)
		if big.NewInt(0).Cmp(&hashInt) == 0 {
			break
		}
	}
}

//增加区块到区块链里面
func (blc *Blockchain) AddBlockToBlockchain(txs []*Transaction) {
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

			newBlock := NewBlock(txs, block.Height+1, block.Hash)
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

func CreateBlockchainWithGenesisBlock(address string) *Blockchain {
	// 判断数据库是否存在
	if dbExists() {
		fmt.Println("创世区块已经存在....")
		os.Exit(2)
	}
	//创建或打开数据库
	db, err := bolt.Open(dbName, 0666, nil)
	if err != nil {
		log.Fatal(err)
	}
	// 关闭数据库

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
			//创建了一个coinbase Transaction
			txCoinBase := NewCoinBaseTransaction(address)

			genesisBlock := CreateGenesisBlock([]*Transaction{txCoinBase})
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
	return &Blockchain{blockHash, db}
}

/**
 * 返回 Blokchain 对象
 */
func BlockChainObject() *Blockchain {

	if !dbExists() {
		log.Println("数据不存在...")
		os.Exit(0)
	}

	//创建或打开数据库
	db, err := bolt.Open(dbName, 0666, nil)
	if err != nil {
		log.Fatal(err)
	}

	var tip []byte
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			tip = b.Get([]byte("l"))
		}
		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	return &Blockchain{tip, db}
}

// 如果一个地址对应的TXOutput 未花费，那么这个Transaction就应该添加到数据中返回
func (blockchain *Blockchain) UnUTXOs(address string, txs []*Transaction) []*UTXOS {

	// 未消费
	var unUTXOs []*UTXOS

	spentTXOutputs := make(map[string][]int)
	// input 消费

	for _, tx := range txs {
		// txHash
		//Vins
		if tx.IsCoinbaseTransactions() == false {

			for _, in := range tx.Vins {
				//是否能够解锁 - 公钥Hash 进行编码
				publicKeyHash := Base58Decode([]byte(address))
				//将第一个 和 最后四个干掉 ,拿到了公钥hash ,进行解锁
				ripemd160Hash := publicKeyHash[1 : len(publicKeyHash)-4]

				if in.UnLockRipemd160Hash(ripemd160Hash) {

					key := hex.EncodeToString(in.TxHash)
					//就有了 input数据
					spentTXOutputs[key] = append(spentTXOutputs[key], in.Vout)
				}
			}
		}

	}

	for _, tx := range txs {

	Work1:
		for index, out := range tx.Vouts {

			for hash, indexArray := range spentTXOutputs {

				txHashStr := hex.EncodeToString(tx.TxHash)

				if hash == txHashStr {

					var isUnSpentUTXO bool

					for _, outIndex := range indexArray {
						if index == outIndex {
							isUnSpentUTXO = true
							continue Work1
						}

						if isUnSpentUTXO == false {
							txout := &UTXOS{tx.TxHash, index, out}
							unUTXOs = append(unUTXOs, txout)
						}
					}

				} else {
					txout := &UTXOS{tx.TxHash, index, out}
					unUTXOs = append(unUTXOs, txout)
				}
			}
		}
	}

	blockIterator := blockchain.Iterator()

	for {

		block := blockIterator.Next()

		for _, tx := range block.Txs {

			// txHash
			//Vins
			if tx.IsCoinbaseTransactions() == false {

				for _, in := range tx.Vins {
					//是否能够解锁 - 公钥Hash 进行编码
					publicKeyHash := Base58Decode([]byte(address))
					//将第一个 和 最后四个干掉 ,拿到了公钥hash ,进行解锁
					ripemd160Hash := publicKeyHash[1 : len(publicKeyHash)-4]
					if in.UnLockRipemd160Hash(ripemd160Hash) {

						key := hex.EncodeToString(in.TxHash)
						//就有了 input数据
						spentTXOutputs[key] = append(spentTXOutputs[key], in.Vout)
					}
				}
			}

			//Vouts
		work:
			for index, out := range tx.Vouts {

				//解锁
				if out.UnLockWithAdress(address) {

					if spentTXOutputs != nil {

						if len(spentTXOutputs) != 0 {

							//判断是谁的hash
							for txHash, indexArray := range spentTXOutputs {

								for _, i := range indexArray {

									if index == i && txHash == hex.EncodeToString(tx.TxHash) {
										continue work
									} else {
										txout := &UTXOS{tx.TxHash, index, out}
										unUTXOs = append(unUTXOs, txout)
									}
								}

							}
						} else {
							txout := &UTXOS{tx.TxHash, index, out}
							unUTXOs = append(unUTXOs, txout)
						}
					}
				}
			}
		}

		var hashInt big.Int
		hashInt.SetBytes(block.PrevBlockHash)

		if hashInt.Cmp(big.NewInt(0)) == 0 {
			break
		}

	}
	return unUTXOs
}

//转账时查找 可用的 UTXO
func (blockchain *Blockchain) FindSpendableUTXOS(from string, amount int64, txs []*Transaction) (int64, map[string][]int) {

	// 1.现获取 所有的 UTXO

	utxos := blockchain.UnUTXOs(from, txs)

	spendableUTXO := make(map[string][]int)
	// 2.遍历 utxos

	var value int64

	for _, utxo := range utxos {

		value = value + utxo.Output.Value

		hash := hex.EncodeToString(utxo.TxHash)
		spendableUTXO[hash] = append(spendableUTXO[hash], utxo.Index)

		if value >= int64(amount) {
			break
		}
	}

	if value < int64(amount) {
		fmt.Printf("%s's fund is 不足\n", from)
		os.Exit(1)
	}
	return value, spendableUTXO
}

// 挖掘新的区块
func (blockchain *Blockchain) MineNewBlock(from, to, amounts []string) {

	//go run main.go send -from '["jack","luo"]' -to '["tom","tea"]' -amount '["10","90"]'

	//1.建立一笔交易

	var txs []*Transaction

	for _, address := range from {
		amount, _ := strconv.Atoi(amounts[0])
		tx := NewSimpleTransaction(address, to[0], int64(amount), blockchain, txs)
		txs = append(txs, tx)
	}

	//1。通过相关算法建立Transaction 数组

	var block *Block

	blockchain.DB.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			hash := b.Get([]byte(blockNewHeader))
			blockBytes := b.Get(hash)
			block = DeserializeBlock(blockBytes)
		}

		return nil
	})

	//2.建立新的区块
	block = NewBlock(txs, block.Height+1, block.Hash)

	// 3.更新存储到数据库

	blockchain.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			b.Put(block.Hash, block.Serialize())
			b.Put([]byte(blockNewHeader), block.Hash)

			blockchain.Tip = block.Hash
		}

		return nil
	})

}

// 查询 余额

func (blockchain *Blockchain) GetBalance(address string) int64 {

	txOutputs := blockchain.UnUTXOs(address, []*Transaction{})
	var amount int64
	for _, utxso := range txOutputs {
		amount += utxso.Output.Value
	}
	return amount
}
