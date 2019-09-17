package BLC

import (
	"fmt"
	"os"
)

//转账
func (cli *Cli) send(from, to, amount []string) {
	if dbExists() == false {
		fmt.Println("数据不存在......")
		os.Exit(1)
	}

	blockchain := BlockChainObject()
	defer blockchain.DB.Close()

	blockchain.MineNewBlock(from, to, amount)
}
