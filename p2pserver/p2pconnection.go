package p2pserver

import (
	"bitbucket.org/blockchain/blockchain"
	"bitbucket.org/blockchain/transaction"
)

// connected clients
var Broadcast = make(chan interface{})

type P2pMessage struct {
	Type        string                  `json:"type"`
	Chain       blockchain.BlockChain   `json:"chain"`
	Transaction transaction.Transaction `json:"transaction"`
}
