package BLC

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const walletFile = "walletFile.dat"

type Wallets struct {
	Walets map[string]*Wallet
}

//  创建钱包集合
func NewWallets() (*Wallets, error) {

	if _, err := os.Stat(walletFile); os.IsNotExist(err) {
		wallets := &Wallets{}
		wallets.Walets = make(map[string]*Wallet)
		wallets.SaveWallets()
	}

	fileContent, err := ioutil.ReadFile(walletFile)
	if err != nil {
		log.Panic(err)
	}
	var wallets Wallets
	gob.Register(elliptic.P256())

	decoder := gob.NewDecoder(bytes.NewReader(fileContent))

	err = decoder.Decode(&wallets)
	if err != nil {
		log.Panic(err)
	}

	return &wallets, nil
}

func (w *Wallets) CreateNewWallet() {
	wallet := NewWallet()
	fmt.Printf("Address: %s\n", wallet.GetAddress())
	w.Walets[string(wallet.GetAddress())] = wallet

	w.SaveWallets()
}

func (w *Wallets) SaveWallets() {
	var content bytes.Buffer

	//注册的目的，是为了，可以序列化任何类型
	gob.Register(elliptic.P256())

	encoder := gob.NewEncoder(&content)

	err := encoder.Encode(&w)
	if err != nil {
		log.Panic(err)
	}
	err = ioutil.WriteFile(walletFile, content.Bytes(), 0644)
	if err != nil {
		log.Panic(err)
	}
}
