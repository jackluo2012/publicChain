package BLC

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	ripemd1602 "golang.org/x/crypto/ripemd160"
	"log"
)

const version = byte(0x00)
const addressCHecksumLen = 4

type Wallet struct {
	// 1.私钥
	PrivateKey ecdsa.PrivateKey
	// 2.公钥(私钥生成的公钥)
	PublicKey []byte
}

//创建钱包
func NewWallet() *Wallet {

	//通过私要
	privateKey, publicKey := newKeyPair()

	return &Wallet{privateKey, publicKey}
}

func IsValidForAddress(address []byte) bool {

	version_public_checksumBytes := Base58Decode(address)

	checkSumBytes := version_public_checksumBytes[len(version_public_checksumBytes)-addressCHecksumLen:]
	version_ripemd160 := version_public_checksumBytes[len(version_public_checksumBytes)-addressCHecksumLen:]

	fmt.Println(len(checkSumBytes))
	fmt.Println(len(version_ripemd160))

	if bytes.Compare(checkSumBytes, version_ripemd160) == 0 {
		return true
	}

	return false
}

func (w *Wallet) GetAddress() []byte {
	// 1. hash160  先将PublicKey
	ripemd160Hash := Ripemd160Hash(w.PublicKey)

	version_ripemd160Hash := append([]byte{version}, ripemd160Hash...)

	checkSumBytes := CheckSum(version_ripemd160Hash)

	bashBytes := append(version_ripemd160Hash, checkSumBytes...)

	return Base58Encode(bashBytes)
}

/**
 * 返回 2次 sha256 后的数组 4位的
 */
func CheckSum(payload []byte) []byte {
	firstSHA := sha256.Sum256(payload)
	secondSHA := sha256.Sum256(firstSHA[:])

	return secondSHA[:addressCHecksumLen]

}

func Ripemd160Hash(publicKey []byte) []byte {

	//1. 256

	hash256 := sha256.New()

	hash256.Write(publicKey)

	hash := hash256.Sum(nil)

	// 2. 160
	ripemd160 := ripemd1602.New()
	ripemd160.Write(hash)

	return ripemd160.Sum(nil)

}

// 通过私钥 产生 公角
func newKeyPair() (ecdsa.PrivateKey, []byte) {

	//1.

	curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

	return *private, pubKey
}
