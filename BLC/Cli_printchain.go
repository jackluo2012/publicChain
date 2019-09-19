package BLC

func (cli *Cli) printchain() {

	blockchain := BlockChainObject()
	//真的要关闭吗？
	defer blockchain.DB.Close()

	blockchain.PrintChain()
}