package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"strconv"
	"time"
)

type Block struct {
	Hash         []byte         `json:"hash"`
	Transactions []*Transaction `json:"transactions"`
	PrevHash     []byte         `json:"previous_hash"`
	Timestamp    int64          `json:"timestamp"`
}

func (b *Block) GetBlockHash() []byte {
	jsonTransactions, _ := json.MarshalIndent(b.Transactions, "", "\t")
	blockInfo := append([]byte(jsonTransactions), b.PrevHash...)
	blockInfo = append(blockInfo, []byte(strconv.Itoa(int(b.Timestamp)))...)

	hash := sha256.Sum256(blockInfo)

	return hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{Hash: []byte{},
		Transactions: []*Transaction{},
		PrevHash:     prevHash,
		Timestamp:    time.Now().Unix(),
	}

	block.Hash = block.GetBlockHash()

	return block
}
