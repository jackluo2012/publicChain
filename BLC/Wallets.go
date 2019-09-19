package BLC

import "fmt"

type Wallets struct {
	Walets map[string]*Wallet
}

//  创建钱包集合
func NewWallets() *Wallets {

	wallets := &Wallets{}
	wallets.Walets = make(map[string]*Wallet)

	return wallets
}

func (w *Wallets) CreateNewWallet() {
	wallet := NewWallet()
	fmt.Printf("Address: %s\n", wallet.GetAddress())
	w.Walets[string(wallet.GetAddress())] = wallet
}
