package transaction

import (
	"bitbucket.org/blockchain/time"
	"bitbucket.org/blockchain/wallet"
	"encoding/json"
	"errors"
	"github.com/pborman/uuid"
)

type output struct {
	Amount  float64 `json:"amount"`
	Address string  `json:"address"`
}

type input struct {
	Amount    float64 `json:"amount"`
	TimeStamp string  `json:"time_stamp"`
	Address   string  `json:"address"`
	Signature string  `json:"signature"`
}

type Transaction struct {
	Id      string   `json:"id"`
	Input   input    `json:"input"`
	Outputs []output `json:"outputs"`
}

func NewTransaction(senderAdd wallet.Wallet, recipient string, amount float64) (Transaction, error) {

	if amount > senderAdd.Balance {
		return Transaction{}, errors.New("insufficient balance")
	}

	outputs := []output{
		{senderAdd.Balance - amount, senderAdd.PublicKey},
		{amount, recipient}}
	transaction := Transaction{uuid.NewRandom().String(), nil, outputs}
	transaction.SignTransaction(senderAdd)
	return transaction, nil
}

func (transaction *Transaction) SignTransaction(senderWallet wallet.Wallet) {

	bytes, _ := json.Marshal(transaction.Outputs)
	transaction.Input = input{senderWallet.Balance, time.MillisInString(), senderWallet.PublicKey, senderWallet.Sign(string(bytes))}
}
