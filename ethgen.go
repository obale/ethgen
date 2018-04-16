package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

func GenerateAddress() (address string, privateKey string, err error) {
	var key *ecdsa.PrivateKey

	if key, err = crypto.GenerateKey(); err != nil {
		return
	}

	address = crypto.PubkeyToAddress(key.PublicKey).Hex()
	privateKey = hex.EncodeToString(key.D.Bytes())
	return
}

func main() {
	var (
		address    string
		privateKey string
		err        error
	)

	if address, privateKey, err = GenerateAddress(); err != nil {
		panic(err)
	}

	fmt.Println(address)
	fmt.Println(privateKey)
}
