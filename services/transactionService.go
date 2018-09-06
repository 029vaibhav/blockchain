package services

import (
	. "bitbucket.org/blockchain/dto"
	. "bitbucket.org/blockchain/transaction"
	"bitbucket.org/blockchain/wallet"
)

var v, _ = wallet.NewWallet(100)

func CreateTransaction(transactionReq TransactionReq) (*Transaction, error) {

	return NewTransaction(&v, transactionReq.RecipientAdd, transactionReq.Amount)

}
