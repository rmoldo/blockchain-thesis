package main

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

type Blockchain struct {
	Blocks []*Block
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

func (chain *Blockchain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)

	chain.Blocks = append(chain.Blocks, newBlock)
}

func InitBlockchain() *Blockchain {
	return &Blockchain{[]*Block{GenerateGenesisBlock()}}
}

func GenerateGenesisBlock() *Block {
	return CreateBlock("Genesis Block", []byte{})
}
