package services

import (
	"bitbucket.org/blockchain/block"
	"bitbucket.org/blockchain/blockchain"
	"bitbucket.org/blockchain/p2pserver"
	"bitbucket.org/blockchain/util"
	log "github.com/sirupsen/logrus"
)

var mainBlock = blockchain.BlockChain{}

func getMainBlockChain() *blockchain.BlockChain {
	return &mainBlock
}

func GetBlockChain() []block.Block {
	return getMainBlockChain().Blocks
}

func AddBlockToBlockChain(data interface{}) block.Block {
	chain := getMainBlockChain()
	addedBlock := chain.AddBlock(data)
	sendBlock(*getMainBlockChain())
	return addedBlock
}

func Replace(newBlockChain blockchain.BlockChain) {
	chain := getMainBlockChain()
	e, b := chain.ReplaceChain(newBlockChain)
	log.Infoln("chain replaced ", b)
	if e == nil && b {
		sendBlock(*getMainBlockChain())
	}

}

func sendBlock(blocks blockchain.BlockChain) {
	log.Infoln("send block message")
	go SendMessage(p2pserver.P2pMessage{Type: util.Blocks, Chain: blocks})
}
