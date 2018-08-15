package blockchain

import "bitbucket.org/blockchain/block"

var blocks []block.Block

func AddBlock(data string) {

	if blocks == nil {
		blocks = append(blocks, block.Genesis())
	}
	minedBlock := block.MineBlock(blocks[len(blocks)-1], data)
	blocks = append(blocks, minedBlock)
}

func GetBlockChain() *[]block.Block {
	return &blocks
}
