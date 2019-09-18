package BLC

import "fmt"

// 查询余额  先用它去 查询 余额
func (cli *Cli) getBalance(address string) {
	fmt.Println("")

	blockchain := BlockChainObject()
	defer blockchain.DB.Close()

	amount := blockchain.GetBalance(address)
	fmt.Println(amount)
}
