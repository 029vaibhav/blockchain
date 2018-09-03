package wallet

import (
	. "crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
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

func NewWallet() (Wallet, error) {

	key, e := GenKeyPair()
	if e != nil {
		return Wallet{}, e
	}
	publicKey := key.PublicKey
	marshal := elliptic.Marshal(publicKey.Curve, publicKey.X, publicKey.Y)
	s := hex.EncodeToString(marshal)
	wallet := Wallet{0, key, s}
	return wallet, nil
}

func (wallet *Wallet) Sign(data string) string {

	r, s, err := Sign(rand.Reader, wallet.KeyPair, []byte(data))
	elliptic.Mar
}
