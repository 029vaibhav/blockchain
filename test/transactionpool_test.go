package test

import (
	"bitbucket.org/blockchain/transaction"
	"bitbucket.org/blockchain/wallet"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransactionPool(t *testing.T) {

	newWallet, _ := wallet.NewWallet()
	newWallet.Balance = 50
	transaction1, _ := transaction.NewTransaction(&newWallet, "abc", 10)
	transaction.UpdateOrAddTransaction(transaction1)
	i := transaction.GetPool().IdTransactionMap[transaction1.Id]
	assert.Equal(t, i.Id, transaction1.Id, "they should be equal")
	var oldTransaction = *transaction1
	newTransaction, _ := transaction1.Update(&newWallet, "some other recipient", 10)
	transaction.UpdateOrAddTransaction(newTransaction)
	i = transaction.GetPool().IdTransactionMap[newTransaction.Id]
	assert.NotEqual(t, i.Id, newTransaction, oldTransaction)

}

func TestCreateTransaction(t *testing.T) {

	newWallet, _ := wallet.NewWallet()
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
