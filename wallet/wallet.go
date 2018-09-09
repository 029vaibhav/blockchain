package wallet

import (
	. "crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"math/big"
)

type Wallet struct {
	Balance   float64    `json:"balance"`
	KeyPair   PrivateKey `json:"key_pair"`
	PublicKey string     `json:"public_key"`
}

func GenKeyPair() (PrivateKey, error) {

	p256 := elliptic.P256()
	privateKey, e := GenerateKey(p256, rand.Reader)
	return *privateKey, e
}

func NewWallet(float642 float64) (*Wallet, error) {

	key, e := GenKeyPair()
	if e != nil {
		return &Wallet{}, e
	}
	publicKey := key.PublicKey
	marshal := elliptic.Marshal(publicKey.Curve, publicKey.X, publicKey.Y)
	s := hex.EncodeToString(marshal)
	wallet := Wallet{float642, key, s}
	return &wallet, nil
}

type Signature struct {
	R *big.Int
	S *big.Int
}

func BlockChainWallet() *Wallet {
	wallet, _ := NewWallet(1000)
	wallet.PublicKey = "block-chain-wallet"
	return wallet
}
