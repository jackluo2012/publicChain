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
	fmt.Println("\tcreateGenesisBlock  -- 输出区块信息.")

}

//检查 参数 是否合法
func isValidArgs() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
}
func (cli *Cli) addBlock(data string) {
	if !dbExists() {
		fmt.Println("数据不存在....")
		os.Exit(1)
	}
	blockchain := BlockChainObject()
	//真的要关闭吗？
	defer blockchain.DB.Close()

	blockchain.AddBlockToBlockchain(data)
}

func (cli *Cli) printchain() {

	blockchain := BlockChainObject()
	//真的要关闭吗？
	defer blockchain.DB.Close()

	blockchain.PrintChain()
}

func (cli *Cli) createGenesisBlockchain(data string) {

	CreateBlockchainWithGenesisBlock(data)

	fmt.Println(data)
}

func (cli *Cli) Run() {
	isValidArgs()

	addBlockFlagCmd := flag.NewFlagSet("addBlock", flag.ExitOnError)
	printChainFlagCmd := flag.NewFlagSet("printChain", flag.ExitOnError)
	createBlockChainCmd := flag.NewFlagSet("createGenesisBlock", flag.ExitOnError)

	flagAddBlockData := addBlockFlagCmd.String("data", "", "")
	flagcreateBlockChainWithData := createBlockChainCmd.String("data", "Genesis Data .....", "")

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
		blockchain.AddBlockToBlockchain(*flagAddBlockData)

	}
	//循环打印
	if printChainFlagCmd.Parsed() {
		blockchain := BlockChainObject()
		defer blockchain.DB.Close()
		blockchain.PrintChain()
	}
	//创建创世区块链
	if createBlockChainCmd.Parsed() {
		cli.createGenesisBlockchain(*flagcreateBlockChainWithData)
	}

}
