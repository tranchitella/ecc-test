package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"fmt"
	"io/ioutil"
)

const (
	privKeyFilename = "secp256k1.der"
	msg             = "hello, world"
)

func decode(data []byte) (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	privateKey, err := x509.ParseECPrivateKey(data)
	if err != nil {
		fmt.Println(err.Error())
	}

	return privateKey, &privateKey.PublicKey
}

func main() {
	priv, err := ioutil.ReadFile(privKeyFilename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	privKey, publicKey := decode(priv)

	hash := sha256.Sum256([]byte(msg))

	r, s, err := ecdsa.Sign(rand.Reader, privKey, hash[:])
	if err != nil {
		panic(err)
	}
	fmt.Printf("signature: (0x%x, 0x%x)\n", r, s)

	valid := ecdsa.Verify(publicKey, hash[:], r, s)
	fmt.Println("signature verified:", valid)
}
