package blockchain

import "time"

type Transaction struct {
	SenderPublicKey   []byte `json:"sender_public_key"`
	ReceiverPublicKey []byte `json:"receiver_public_key"`
	Data              string `json:"data"`
	Type              string `json:"type"`
	Timestamp         int64  `json:"timestamp"`
}

func CreateTransaction(senderPublicKey []byte, receiverPublicKey []byte, data string, transactionType string) *Transaction {
	transaction := &Transaction{SenderPublicKey: senderPublicKey,
		ReceiverPublicKey: receiverPublicKey,
		Data:              data,
		Type:              transactionType,
		Timestamp:         time.Now().Unix(),
	}

	return transaction
}
