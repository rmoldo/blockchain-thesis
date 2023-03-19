package main

import (
	"encoding/json"
	"fmt"
	"thesis/blockchain"
)

func main() {
	chain := blockchain.InitBlockchain()

	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")
	chain.AddBlock("Fourth Block")

	for _, block := range chain.Blocks {
		jsonBlock, _ := json.MarshalIndent(block, "", "\t")
		fmt.Println(string(jsonBlock))

	}

	wallet := blockchain.InitWallet()

	transaction := wallet.CreateTransaction(wallet.PublicKey.N.Bytes(), "Cereale", "TRANSFER")
	jsonTransaction, _ := json.MarshalIndent(transaction, "", "\t")
	fmt.Println(string(jsonTransaction))

	tmpTrans, _ := json.MarshalIndent(transaction.GetTransactionWithDefaultSignature(), "", "\t")

	fmt.Println(wallet.VerifySignature([]byte(tmpTrans), transaction.Signature, wallet.PublicKey))
}
