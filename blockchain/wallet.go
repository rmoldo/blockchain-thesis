package blockchain

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
)

type Wallet struct {
	PrivateKey rsa.PrivateKey
	PublicKey  []byte
}
