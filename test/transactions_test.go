package test

import (
	"bitbucket.org/blockchain/transaction"
	"bitbucket.org/blockchain/wallet"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestTransaction(t *testing.T) {

	newWallet, _ := wallet.NewWallet()
	newWallet.Balance = 50
	newTransaction, _ := transaction.NewTransaction(&newWallet, "abc", 10)
	assert.Equal(t, newWallet.Balance, newTransaction.Input.Amount)

}

func TestSignatureVerification(t *testing.T) {

	newWallet, _ := wallet.NewWallet()
	newWallet.Balance = 50
	newTransaction, _ := transaction.NewTransaction(&newWallet, "abc", 10)
	assert.Equal(t, newTransaction.VerifyTransaction(), true)

}

func TestCorruptSignatureVerification(t *testing.T) {

	newWallet, _ := wallet.NewWallet()
	newWallet.Balance = 50
	newTransaction, _ := transaction.NewTransaction(&newWallet, "abc", 10)
	newTransaction.Outputs[0].Amount = 5000
	assert.Equal(t, newTransaction.VerifyTransaction(), false)

}

func TestUpdateTransaction(t *testing.T) {

	newWallet, _ := wallet.NewWallet()
	newWallet.Balance = 50
	newTransaction, _ := transaction.NewTransaction(&newWallet, "abc", 10)
	update, _ := newTransaction.Update(&newWallet, "abc", 20)
	var o transaction.Output
	for _, v := range update.Outputs {
		if v.Address == newWallet.PublicKey {
			o = v
		}
	}
	assert.Equal(t, o.Amount, newWallet.Balance-10-20)
	for _, v := range update.Outputs {
		if v.Address == "abc" {
			o = v
		}
	}
	assert.Equal(t, o.Amount, float64(20))

}
