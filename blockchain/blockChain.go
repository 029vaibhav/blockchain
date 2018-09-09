package blockchain

import (
	"bitbucket.org/blockchain/block"
	"bitbucket.org/blockchain/transaction"
	"errors"
	"github.com/google/go-cmp/cmp"
	log "github.com/sirupsen/logrus"
	"gopkg.in/square/go-jose.v1/json"
)

type BlockChain struct {
	Blocks []block.Block `json:"blocks"`
}

func (b *BlockChain) AddBlock(data interface{}) block.Block {

	if b.Blocks == nil {
		b.Blocks = append(b.Blocks, block.Genesis())
	}
	minedBlock := block.MineBlock(b.Blocks[len(b.Blocks)-1], data)
	b.Blocks = append(b.Blocks, minedBlock)
	return minedBlock
}

func (b *BlockChain) IsValidChain(chain BlockChain) bool {

	if len(b.Blocks) == 0 {
		return true
	}

	if !cmp.Equal(chain.Blocks[0], b.Blocks[0]) {
		return false
	}
	for i := 1; i < len(chain.Blocks); i++ {
		x := chain.Blocks[i]
		currentBlock := &x
		lastBlock := chain.Blocks[i-1]

		if currentBlock.LastHash != lastBlock.Hash || currentBlock.Hash != block.GetBlockHash(*currentBlock) {
			return false
		}
		transactions := make([]transaction.Transaction, 0)
		bytes, _ := json.Marshal(currentBlock.Data)
		currentBlock.Data = json.Unmarshal(bytes, transactions)

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
	b = &chain
	return nil, true
}
