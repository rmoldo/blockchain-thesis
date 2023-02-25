package blockchain

import (
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Hash      []byte `json:"hash"`
	Data      string `json:"data"` // TODO: change to transactions when implemented
	PrevHash  []byte `json:"previous_hash"`
	Timestamp int64  `json:"timestamp"`
}

func (b *Block) GetBlockHash() []byte {
	blockInfo := append([]byte(b.Data), b.PrevHash...)
	blockInfo = append(blockInfo, []byte(strconv.Itoa(int(b.Timestamp)))...)

	hash := sha256.Sum256(blockInfo)

	return hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{Hash: []byte{},
		Data:      data,
		PrevHash:  prevHash,
		Timestamp: time.Now().Unix(),
	}

	block.Hash = block.GetBlockHash()

	return block
}
