package main

import (
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Hash      []byte
	Data      string // TODO: to be changed when the time comes
	PrevHash  []byte
	Timestamp int64
}

// Compute the hash of a block and return it as a slice of bytes
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
