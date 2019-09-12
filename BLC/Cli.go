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
	fmt.Println("\taddBlock -data DATA -- 交易数据.")
	fmt.Println("\tprintChain  -- 输出区块信息.")
	fmt.Println("\tcreateGenesisBlock -address  -- 输出区块信息.")

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

func (cli *Cli) printchain() {

	blockchain := BlockChainObject()
	//真的要关闭吗？
	defer blockchain.DB.Close()

	blockchain.PrintChain()
}

func (cli *Cli) createGenesisBlockchain(address string) {


	// 1. 创建coinBase

	CreateBlockchainWithGenesisBlock(address)

	fmt.Println(address)
}

func (cli *Cli) Run() {
	isValidArgs()

	addBlockFlagCmd := flag.NewFlagSet("addBlock", flag.ExitOnError)
	printChainFlagCmd := flag.NewFlagSet("printChain", flag.ExitOnError)
	createBlockChainCmd := flag.NewFlagSet("createGenesisBlock", flag.ExitOnError)

	flagAddBlockData := addBlockFlagCmd.String("data", "", "")
	flagcreateBlockChainWithAddress := createBlockChainCmd.String("address", "", "创建创世区块的地址")

	switch os.Args[1] {
	case "addBlock":
		err := addBlockFlagCmd.Parse(os.Args[2:])
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
	default:
		printUsage()
		os.Exit(0)
	}

	//能解析成功
	if addBlockFlagCmd.Parsed() {
		if *flagAddBlockData == "" {
			printUsage()
			os.Exit(1)
		}
		blockchain := BlockChainObject()
		defer blockchain.DB.Close()
		blockchain.AddBlockToBlockchain([]*Transaction{})

	}
	//循环打印
	if printChainFlagCmd.Parsed() {
		blockchain := BlockChainObject()
		defer blockchain.DB.Close()
		blockchain.PrintChain()
	}
	//创建创世区块链
	if createBlockChainCmd.Parsed() {
		if *flagcreateBlockChainWithAddress ==""{
			fmt.Println("地址不能为空.....")
			printUsage()
			os.Exit(1)
		}

		cli.createGenesisBlockchain(*flagcreateBlockChainWithAddress)
	}

}
