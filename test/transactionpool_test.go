package test

import (
	"bitbucket.org/blockchain/dto"
	"bitbucket.org/blockchain/services"
	"bitbucket.org/blockchain/transaction"
	"bitbucket.org/blockchain/wallet"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransactionPool(t *testing.T) {

	WalletAddress := services.WalletAddress

	transaction1, _ := services.CreateTransaction(dto.TransactionReq{RecipientAdd: "abc", Amount: 10})
	i := transaction.GetTransactions()[transaction1.Id]
	assert.Equal(t, i.Id, transaction1.Id, "they should be equal")
	var oldTransaction = *transaction1
	newTransaction, _ := transaction1.Update(&WalletAddress, "some other recipient", 10)
	transaction.UpdateOrAddTransaction(newTransaction)
	i = transaction.GetTransactions()[newTransaction.Id]
	assert.NotEqual(t, i.Id, newTransaction, oldTransaction)

}

func TestCreateTransaction(t *testing.T) {

	newWallet, _ := wallet.NewWallet(50)
	newWallet.Balance = 50
	recep := "random-address"
	amount := float64(10)
	transaction.NewTransaction(&newWallet, recep, amount)
	newTransaction, _ := transaction.NewTransaction(&newWallet, recep, amount)

	var out transaction.Output
	for _, v := range newTransaction.Outputs {
		if v.Address == newWallet.PublicKey {
			out = v
			break
		}
	}
	assert.Equal(t, out.Amount, newWallet.Balance-amount-amount)

}
