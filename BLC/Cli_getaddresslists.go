package BLC

import "fmt"

//打印所有的钱包地址

func (cli *Cli) addressLists() {

	fmt.Println("打印所有的钱包地址:")
	wallets, _ := NewWallets()

	for address, _ := range wallets.Walets {
		fmt.Println(address)
	}
}
