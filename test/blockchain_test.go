package test

import (
	"bytes"
	"testing"
	"thesis/blockchain"
)

func TestAddBlock(t *testing.T) {
	chain := blockchain.InitBlockchain()

	chain.AddBlock("First Block")
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")
	
	if !bytes.Equal(chain.Blocks[0].Hash, chain.Blocks[1].PrevHash) && !bytes.Equal(chain.Blocks[1].Hash, chain.Blocks[2].PrevHash) && !bytes.Equal(chain.Blocks[2].Hash, chain.Blocks[3].PrevHash) {
		t.Errorf("Expected chain to be completed")
	}
}

func TestBlockchainData(t *testing.T) {
	chain := blockchain.InitBlockchain()

	chain.AddBlock("First Block")
	chain.AddBlock("Second Block")
	
	if chain.Blocks[1].Data != "First Block" || chain.Blocks[2].Data != "Second Block" {
		t.Errorf("Data in block")
	}
}
