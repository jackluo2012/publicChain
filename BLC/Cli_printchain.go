package BLC

import "fmt"
//内容
func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\tsendBlock -from FROM -to TO -amount AMOUNT -- 交易明细.")
	fmt.Println("\tprintChain  -- 输出区块信息.")
	fmt.Println("\tcreateGenesisBlock -address  -- 输出区块信息.")
	fmt.Println("\tgetbalance -address  -- 获取帐户余额.")

}



func (cli *Cli) printchain() {

	blockchain := BlockChainObject()
	//真的要关闭吗？
	defer blockchain.DB.Close()

	blockchain.PrintChain()
}