package BLC

import "fmt"

func (cli *Cli) createWallet() {

	wallets, _ := NewWallets()
	wallets.CreateNewWallet()
	fmt.Println(wallets.Walets)

	wallets.SaveWallets()

}
