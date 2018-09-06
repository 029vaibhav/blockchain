package transaction

import (
	"bitbucket.org/blockchain/time"
	"bitbucket.org/blockchain/util"
	"bitbucket.org/blockchain/wallet"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"
)

type Output struct {
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
	Outputs []Output `json:"outputs"`
}

func (transaction *Transaction) Update(senderAdd *wallet.Wallet, recipient string, amount float64) (*Transaction, error) {

	i := 0
	for _, v := range transaction.Outputs {
		if v.Address == senderAdd.PublicKey {
			break
		}
		i++
	}
	output := &transaction.Outputs[i]
	if amount > output.Amount {
		return nil, errors.New("insufficient balance")
	}
	output.Amount = output.Amount - amount
	transaction.Outputs = append(transaction.Outputs, Output{amount, recipient})
	transaction.SignTransaction(senderAdd)
	return transaction, nil

}

func NewTransaction(senderAdd *wallet.Wallet, recipient string, amount float64) (*Transaction, error) {

	if amount > senderAdd.Balance {
		return nil, errors.New("insufficient balance")
	}

	existingTransaction := GetExistingTransaction(senderAdd.PublicKey)
	if existingTransaction == nil {
		outputs := []Output{
			{senderAdd.Balance - amount, senderAdd.PublicKey},
			{amount, recipient}}
		existingTransaction = &Transaction{util.RandomIdGen(), input{}, outputs}
		existingTransaction.SignTransaction(senderAdd)
		UpdateOrAddTransaction(existingTransaction)

	} else {
		existingTransaction.Update(senderAdd, recipient, amount)
	}
	return existingTransaction, nil
}

func (transaction *Transaction) SignTransaction(senderWallet *wallet.Wallet) error {

	i := transaction.getOutputAsString()
	s, e := sign(util.GetHash(i), &senderWallet.KeyPair)
	if e == nil {
		transaction.Input = input{senderWallet.Balance, time.MillisInString(), senderWallet.PublicKey, s}
		return nil
	}
	return e
}

func sign(data string, KeyPair *ecdsa.PrivateKey) (string, error) {

	r, s, err := ecdsa.Sign(rand.Reader, KeyPair, []byte(data))
	if err == nil {
		signature := r.Bytes()
		signature = append(signature, s.Bytes()...)
		return hex.EncodeToString(signature), nil
	}
	return "", err
}

func (transaction *Transaction) getOutputAsString() string {
	var builder strings.Builder
	for _, v := range transaction.Outputs {
		builder.WriteString(v.Address)
		builder.WriteString(fmt.Sprintf("%f", v.Amount))
	}
	i := builder.String()
	return i
}

func (transaction *Transaction) VerifyTransaction() bool {
	return VerifySignature(transaction.Input.Address, transaction.Input.Signature, util.GetHash(transaction.getOutputAsString()))
}

func VerifySignature(publicKey string, signature string, dataHash string) bool {

	bytes := []byte(dataHash)

	p256 := elliptic.P256()
	decodeString, _ := hex.DecodeString(publicKey)
	x, y := elliptic.Unmarshal(p256, decodeString)
	key := getPubKey(x, y)
	decodedSign, _ := hex.DecodeString(signature)
	var r []byte
	for i := 0; i < 32; i++ {
		r = append(r, decodedSign[i])
	}
	var s []byte
	for i := 32; i < 64; i++ {
		s = append(s, decodedSign[i])
	}

	setBytesR := new(big.Int).SetBytes(r)
	setBytesS := new(big.Int).SetBytes(s)
	return ecdsa.Verify(key, bytes, setBytesR, setBytesS)
}

func getPubKey(x, y *big.Int) *ecdsa.PublicKey {

	pub := new(ecdsa.PublicKey)
	pub.X = x
	pub.Y = y

	pub.Curve = elliptic.P256()
	return pub
}
