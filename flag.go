package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)
//内容
func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\taddBlock -data DATA -- 交易数据.")
	fmt.Println("\tprintChain  -- 输出区块信息.")

}
//检查 参数 是否合法
func isValidArgs() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
}

func main() {

	isValidArgs()

	addBlockFlagCmd := flag.NewFlagSet("addBlock", flag.ExitOnError)
	printChainFlagCmd := flag.NewFlagSet("printChain", flag.ExitOnError)

	flagAddBlockData := addBlockFlagCmd.String("data", "", "")
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
		fmt.Println(*flagAddBlockData)
	}
	if printChainFlagCmd.Parsed() {
		fmt.Println("打印区块的值")
	}

}
