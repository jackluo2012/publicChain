package BLC

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type Cli struct {
}

//内容
func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\tsendBlock -from FROM -to TO -amount AMOUNT -- 交易明细.")
	fmt.Println("\tcreatewallet  -- 创建钱包.")
	fmt.Println("\tprintChain  -- 输出区块信息.")
	fmt.Println("\tcreateGenesisBlock -address  -- 输出区块信息.")
	fmt.Println("\tgetbalance -address  -- 获取帐户余额.")

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

	createWalletCmd := flag.NewFlagSet("createwallet", flag.ExitOnError)

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
	case "createwallet":
		err := createWalletCmd.Parse(os.Args[2:])
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

		for index, fromAddress := range from {

			if IsValidForAddress([]byte(fromAddress)) == false || IsValidForAddress([]byte(to[index])) == false {
				fmt.Println("地址无效....")
				os.Exit(1)
			}
		}

		amount := JSONToArray(*flagAmount)
		cli.send(from, to, amount)
	}
	//循环打印
	if printChainFlagCmd.Parsed() {
		blockchain := BlockChainObject()
		defer blockchain.DB.Close()
		blockchain.PrintChain()
	}

	if createWalletCmd.Parsed() {
		// 创建钱包
		cli.createWallet()
	}

	//创建创世区块链
	if createBlockChainCmd.Parsed() {
		if IsValidForAddress([]byte(*flagcreateBlockChainWithAddress)) == false {
			fmt.Println("地址无效.....")
			printUsage()
			os.Exit(1)
		}

		cli.createGenesisBlockchain(*flagcreateBlockChainWithAddress)
	}

	//创建创世区块链
	if getbalanceCmd.Parsed() {
		if IsValidForAddress([]byte(*getbalanceWithAddress)) == false {
			fmt.Println("地址无效.....")
			printUsage()
			os.Exit(1)
		}

		cli.getBalance(*getbalanceWithAddress)
	}

}
