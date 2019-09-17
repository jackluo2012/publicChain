package BLC

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type Cli struct {
}



//检查 参数 是否合法
func isValidArgs() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
}
func (cli *Cli) addBlock(txs []*Transaction) {
	if !dbExists() {
		fmt.Println("数据不存在....")
		os.Exit(1)
	}
	blockchain := BlockChainObject()
	//真的要关闭吗？
	defer blockchain.DB.Close()

	blockchain.AddBlockToBlockchain([]*Transaction{})
}





func (cli *Cli) Run() {
	isValidArgs()

	sendBlockFlagCmd := flag.NewFlagSet("send", flag.ExitOnError)
	printChainFlagCmd := flag.NewFlagSet("printChain", flag.ExitOnError)
	createBlockChainCmd := flag.NewFlagSet("createGenesisBlock", flag.ExitOnError)
	getbalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)

	flagFrom := sendBlockFlagCmd.String("from", "", "转账原地址......")
	flagTo := sendBlockFlagCmd.String("to", "", "转账目的地地址......")
	flagAmount := sendBlockFlagCmd.String("amount", "", "转账的金额......")

	flagcreateBlockChainWithAddress := createBlockChainCmd.String("address", "", "创建创世区块的地址")
	getbalanceWithAddress := getbalanceCmd.String("address", "", "要查询某一个帐户余额")

	switch os.Args[1] {
	case "send":
		err := sendBlockFlagCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printChain":
		err := printChainFlagCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "createGenesisBlock":
		err := createBlockChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "getbalance":
		err := getbalanceCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		printUsage()
		os.Exit(0)
	}

	//能解析成功
	if sendBlockFlagCmd.Parsed() {
		if *flagFrom == "" || *flagTo == "" || *flagAmount == "" {
			printUsage()
			os.Exit(1)
		}

		from := JSONToArray(*flagFrom)
		to := JSONToArray(*flagTo)
		amount := JSONToArray(*flagAmount)
		cli.send(from, to, amount)
	}
	//循环打印
	if printChainFlagCmd.Parsed() {
		blockchain := BlockChainObject()
		defer blockchain.DB.Close()
		blockchain.PrintChain()
	}
	//创建创世区块链
	if createBlockChainCmd.Parsed() {
		if *flagcreateBlockChainWithAddress == "" {
			fmt.Println("地址不能为空.....")
			printUsage()
			os.Exit(1)
		}

		cli.createGenesisBlockchain(*flagcreateBlockChainWithAddress)
	}

	//创建创世区块链
	if getbalanceCmd.Parsed() {
		if *getbalanceWithAddress == "" {
			fmt.Println("地址不能为空.....")
			printUsage()
			os.Exit(1)
		}

		cli.getBalance(*getbalanceWithAddress)
	}

}
