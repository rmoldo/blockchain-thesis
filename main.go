package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	chain := InitBlockchain()

	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")
	chain.AddBlock("Fourth Block")

	for _, block := range chain.Blocks {
		jsonBlock, _ := json.MarshalIndent(block, "", "\t")
		fmt.Println(string(jsonBlock))

	}
}
