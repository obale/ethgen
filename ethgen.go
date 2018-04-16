package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	var (
		err error
		key *ecdsa.PrivateKey
	)

	if key, err = crypto.GenerateKey(); err != nil {
		panic(err)
	}

	address := crypto.PubkeyToAddress(key.PublicKey).Hex()
	fmt.Println(address)

	privateKey := hex.EncodeToString(key.D.Bytes())
	fmt.Println(privateKey)
}
