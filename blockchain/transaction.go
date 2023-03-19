package blockchain

import "time"

type Transaction struct {
	SenderPublicKey   []byte `json:"sender_public_key"`
	ReceiverPublicKey []byte `json:"receiver_public_key"`
	Data              string `json:"data"`
	Type              string `json:"type"`
	Timestamp         int64  `json:"timestamp"`
	Signature         string `json:"signature"`
}

func CreateTransaction(senderPublicKey []byte, receiverPublicKey []byte, data string, transactionType string) *Transaction {
	transaction := &Transaction{SenderPublicKey: senderPublicKey,
		ReceiverPublicKey: receiverPublicKey,
		Data:              data,
		Type:              transactionType,
		Timestamp:         time.Now().Unix(),
		Signature:         "",
	}

	return transaction
}

func (t *Transaction) SetSignature(signature string) {
	t.Signature = signature
}

func (t *Transaction) GetTransactionWithDefaultSignature() *Transaction {
	transaction := *t
	transaction.Signature = ""

	return &transaction
}
