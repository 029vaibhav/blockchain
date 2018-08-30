package blockchain

import (
	"bitbucket.org/blockchain/block"
	"errors"
	"github.com/google/go-cmp/cmp"
	log "github.com/sirupsen/logrus"
)

type BlockChain struct {
	Blocks []block.Block `json:"blocks"`
}

func (b *BlockChain) AddBlock(data string) {

	if b.Blocks == nil {
		b.Blocks = append(b.Blocks, block.Genesis())
	}
	minedBlock := block.MineBlock(b.Blocks[len(b.Blocks)-1], data)
	b.Blocks = append(b.Blocks, minedBlock)
}

func (b *BlockChain) IsValidChain(chain BlockChain) bool {

	if len(b.Blocks) == 0 {
		return true
	}

	if !cmp.Equal(chain.Blocks[0], b.Blocks[0]) {
		return false
	}
	for i := 1; i < len(chain.Blocks); i++ {
		currentBlock := chain.Blocks[i]
		lastBlock := chain.Blocks[i-1]

		if currentBlock.LastHash != lastBlock.Hash || currentBlock.Hash != block.GetBlockHash(currentBlock) {
			return false
		}
	}
	return true
}

func (b *BlockChain) ReplaceChain(chain BlockChain) (error, bool) {

	if len(chain.Blocks) <= len(b.Blocks) {
		return errors.New("new chain is smaller or equal then the old chain"), false
	} else if !b.IsValidChain(chain) {
		return errors.New("its an invalid chain"), false
	}
	log.Info("replacement of chain completed")
	*b = chain
	return nil, true
}
