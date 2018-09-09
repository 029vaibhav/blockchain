package services

import (
	. "bitbucket.org/blockchain/dto"
	"bitbucket.org/blockchain/p2pserver"
	. "bitbucket.org/blockchain/transaction"
	"bitbucket.org/blockchain/util"
	"bitbucket.org/blockchain/wallet"
	. "strconv"
)

var WalletAddress, _ = wallet.NewWallet(100)

func CreateTransaction(transactionReq TransactionReq) (*Transaction, error) {

	WalletAddress.Balance = CalculateBalance()
	transaction, e := NewTransaction(WalletAddress, transactionReq.RecipientAdd, transactionReq.Amount)
	if e == nil {
		sendTransaction(*transaction, util.Transactions)
	}
	return transaction, e

}

func sendTransaction(transaction Transaction, transactionType string) {
	go SendMessage(p2pserver.P2pMessage{Type: transactionType, Transaction: transaction})
}

func GetWallet() *wallet.Wallet {
	return WalletAddress
}

func ClearTransaction() {
	Clear()
	sendTransaction(Transaction{}, util.ClearTransaction)
}

func CalculateBalance() float64 {

	var latestTransaction *Transaction
	startTime := 0
	balance := WalletAddress.Balance
	var transactions []BalanceCalc

	for _, v := range GetBlockChain() {

		if v.Data == nil {
			continue
		}
		transactionList := v.Data.([]Transaction)
		for _, v := range transactionList {
			if v.Input.Address == WalletAddress.PublicKey {
				if latestTransaction == nil {
					latestTransaction = &v
				} else {
					currentTransactionTS, _ := Atoi(v.Input.TimeStamp)
					latestTransactionTS, _ := Atoi(latestTransaction.Input.TimeStamp)
					if currentTransactionTS > latestTransactionTS {
						latestTransaction = &v
					}
				}
			}
			calc := BalanceCalc{}
			calc.TimeStamp, _ = Atoi(v.Input.TimeStamp)
			var outputs []Output
			for _, output := range v.Outputs {
				if output.Address == WalletAddress.PublicKey {
					outputs = append(outputs, output)
				}
			}
			calc.Outputs = outputs
			transactions = append(transactions, calc)

		}
	}

	if latestTransaction != nil {
		for _, v := range latestTransaction.Outputs {
			if v.Address == WalletAddress.PublicKey {
				balance = v.Amount
				startTime, _ = Atoi(latestTransaction.Input.TimeStamp)
			}
		}
	}

	for _, v := range transactions {
		if v.TimeStamp > startTime {
			for _, v := range v.Outputs {
				balance += v.Amount
			}

		}
	}
	return balance

}
