package main

import (
	"crypto/sha256"
	"fmt"
	ripemd1602 "golang.org/x/crypto/ripemd160"
)

func main() {

	hash265()
	ripemd160()
}

func ripemd160() {


	// 160

	// bit(位) 160  20个 字节 （1字节 8 位）

	hasher := ripemd1602.New()

	hasher.Write([]byte("http://www.google.com"))

	hashBytes := hasher.Sum(nil)

	hashString := fmt.Sprintf("%x", hashBytes)

	fmt.Println(hashString)
}
/**
 * 2个数字为8位
 */
func hash265() {
	hasher := sha256.New()
	hasher.Write([]byte("http://www.google.com"))
	hashBytes := sha256.Sum256(nil)
	hashString := fmt.Sprintf("%x", hashBytes)

	fmt.Println(hashString)
}
