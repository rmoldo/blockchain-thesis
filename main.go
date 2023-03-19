package main

import (
	"encoding/json"
	"fmt"
	"thesis/blockchain"
)

func main() {
	wallet := blockchain.InitWallet()

	transaction := wallet.CreateTransaction(wallet.PublicKey.N.Bytes(), "Cereale", "TRANSFER")
	jsonTransaction, _ := json.MarshalIndent(transaction, "", "\t")
	fmt.Println(string(jsonTransaction))

	tmpTrans, _ := json.MarshalIndent(transaction.GetTransactionWithDefaultSignature(), "", "\t")

	fmt.Println(wallet.VerifySignature([]byte(tmpTrans), transaction.Signature, wallet.PublicKey))
}
