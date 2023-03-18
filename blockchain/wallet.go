package blockchain

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

type Wallet struct {
	PrivateKey rsa.PrivateKey
	PublicKey  rsa.PublicKey
}

func (w *Wallet) InitKeys() {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 4096)
	w.PrivateKey = *privateKey

	w.PublicKey = privateKey.PublicKey
}

func (w *Wallet) Sign(data []byte) []byte {
	// Hash the data
	dataHash := sha256.New()
	_, err := dataHash.Write(data)

	if err != nil {
		// TODO: error handling
	}

	dataHashSum := dataHash.Sum(nil)

	signature, err := rsa.SignPSS(rand.Reader, &w.PrivateKey, crypto.SHA256, dataHashSum, nil)

	if err != nil {
		// TODO: error handling
	}

	return signature
}
