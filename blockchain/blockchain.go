package blockchain

type Blockchain struct {
	Blocks []*Block
}

// Compute the hash of a block and return it as a slice of bytes
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
