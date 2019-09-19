package BLC

import "fmt"

func (cli *Cli) createWallet() {

	wallets := NewWallets()
	wallets.CreateNewWallet()
	fmt.Println(wallets.Walets)

}
