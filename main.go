package main

import (
	"encoding/hex"
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

	transaction := blockchain.CreateTransaction([]byte("terrance"), []byte("butthead"), "Manele", "Smardeala")

	jsonTransaction, _ := json.MarshalIndent(transaction, "", "\t")

	wallet := blockchain.InitWallet()

	rawSignature := wallet.Sign([]byte(jsonTransaction))

	signature := hex.EncodeToString(rawSignature)

	fmt.Println(wallet.VerifySignature([]byte(jsonTransaction), signature, wallet.PublicKey))

	transaction.SetSignature(hex.EncodeToString(wallet.Sign([]byte(jsonTransaction))))

	jsonTransaction, _ = json.MarshalIndent(transaction, "", "\t")

	fmt.Println(string(jsonTransaction))
}
