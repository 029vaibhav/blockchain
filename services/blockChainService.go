package services

import (
	"bitbucket.org/blockchain/block"
	"bitbucket.org/blockchain/blockchain"
	log "github.com/sirupsen/logrus"
)

var mainBlock = blockchain.BlockChain{}

func getMainBlockChain() *blockchain.BlockChain {
	return &mainBlock
}

func GetBlockChain() []block.Block {
	return getMainBlockChain().Blocks
}

func AddBlockChain(block block.Block) {
	chain := getMainBlockChain()
	chain.AddBlock(block.Data)
	sendMessage()
}

func Replace(newBlockChain blockchain.BlockChain) {
	chain := getMainBlockChain()
	e, b := chain.ReplaceChain(newBlockChain)
	log.Infoln("chain replaced ", b)
	if e == nil && b {
		sendMessage()
	}

}

func sendMessage() {
	go SendMessage(*getMainBlockChain())
}
