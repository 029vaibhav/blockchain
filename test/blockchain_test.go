package test

import "testing"
import (
	"bitbucket.org/blockchain/block"
	"bitbucket.org/blockchain/blockchain"
	"bitbucket.org/blockchain/environment"
	"github.com/stretchr/testify/assert"
)

func TestGenesis(t *testing.T) {

	chain, _ := setup()
	assert.Equal(t, chain.Blocks[0], block.Genesis(), "they should be equal")

}

func TestData(t *testing.T) {

	chain, _ := setup()
	//mentioned data and data in the chain should be same
	data := "2nd block"
	chain.AddBlock(data)
	assert.Equal(t, chain.Blocks[len(chain.Blocks)-1].Data, data, "they should be equal")

}
func TestValidateChain(t *testing.T) {

	chain, chain2 := setup()
	data := "2nd block"
	chain.AddBlock(data)
	assert.True(t, chain.IsValidChain(chain2), "it should be true")

}
func TestInValidateCorruptGenesisBlock(t *testing.T) {

	chain, chain2 := setup()
	//mentioned data and data in the chain should be same
	data := "2nd block"
	chain.AddBlock(data)
	//Invalidates with corrupt genesis block
	chain2.Blocks[0].Data = "bad data"
	assert.False(t, chain.IsValidChain(chain2), "it should be false")

}
func TestInValidateChain(t *testing.T) {

	chain, chain2 := setup()
	data := "2nd block"
	chain.AddBlock(data)
	//Invalidate corrupt chain
	chain2.Blocks[1].Data = " data"
	assert.False(t, chain.IsValidChain(chain2), "it should be false")

}
func TestReplaceValidChain(t *testing.T) {

	chain, chain2 := setup()

	//Replace chain with valid chain
	chain2.AddBlock("3 block")
	chain.ReplaceChain(chain2)
	assert.Equal(t, chain, chain2, "they should be equal")

}

func TestReplaceInValidChain(t *testing.T) {

	chain, chain2 := setup()

	//Replace chain with valid chain
	replaceChain, _ := chain.ReplaceChain(chain2)
	assert.NotNil(t, replaceChain, "it should contain error")

}

func TestHashFunction(t *testing.T) {

	chain, _ := setup()
	substring := chain.Blocks[1].Hash[0:chain.Blocks[1].Difficulty]
	assert.Equal(t, substring, chain.Blocks[1].Difficulty, "they should be equal")

}
func TestDifficultyTest(t *testing.T) {

	chain, _ := setup()
	difficulty := block.AdjustDifficulty(chain.Blocks[1], chain.Blocks[1].Timestamp+"00", 1000)
	assert.Equal(t, difficulty, chain.Blocks[1].Difficulty-1, "they should be equal")

	difficulty = block.AdjustDifficulty(chain.Blocks[1], chain.Blocks[1].Timestamp, 1000)
	assert.Equal(t, difficulty, chain.Blocks[1].Difficulty+1, "they should be equal")

}

func setup() (blockchain.BlockChain, blockchain.BlockChain) {

	environment.Instance().Set("difficulty", 4)
	environment.Instance().Set("mine_rate_in_ms", 1000)

	chain := blockchain.BlockChain{}
	chain.AddBlock("1st block")

	chain2 := blockchain.BlockChain{}
	chain2.AddBlock("1st block")
	return chain, chain2
}
