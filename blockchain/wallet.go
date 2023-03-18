package blockchain

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
)

type Wallet struct {
	PrivateKey rsa.PrivateKey
	PublicKey  rsa.PublicKey
}

func InitWallet() *Wallet {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)

	return &Wallet{PrivateKey: *privateKey,
		PublicKey: privateKey.PublicKey,
	}
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

func (w *Wallet) VerifySignature(data []byte, signature string, publicKey rsa.PublicKey) bool {
	// Hash the data
	dataHash := sha256.New()
	_, err := dataHash.Write(data)

	if err != nil {
		// TODO: error handling
	}

	dataHashSum := dataHash.Sum(nil)

	rawSignature, _ := hex.DecodeString(signature)

	err = rsa.VerifyPSS(&publicKey, crypto.SHA256, dataHashSum, rawSignature, nil)

	if err != nil {
		return false
	}

	return true
}
