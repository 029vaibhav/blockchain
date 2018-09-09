package block

import (
	"bitbucket.org/blockchain/environment"
	"bitbucket.org/blockchain/time"
	"bitbucket.org/blockchain/util"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"strconv"
	"strings"
)

type Block struct {
	Timestamp  string      `json:"timestamp"`
	LastHash   string      `json:"last_hash"`
	Hash       string      `json:"hash"`
	Data       interface{} `json:"data"`
	Nonce      int         `json:"nonce"`
	Difficulty int         `json:"difficulty"`
}

func Genesis() Block {
	block := Block{"Genesis TimeStamp", "-----", "GenesisHash", nil, 0, environment.Instance().Get("difficulty").(int)}
	return block
}

func MineBlock(lastBlock Block, data interface{}) Block {

	difficulty := environment.Instance().Get("difficulty").(int)
	mineRate := environment.Instance().Get("mine_rate_in_ms").(int)
	currentTime := ""
	currentHash := ""
	nonce := -1
	for {
		nonce++
		currentTime = time.MillisInString()
		difficulty = AdjustDifficulty(lastBlock, currentTime, mineRate)
		currentHash = getHash(currentTime, lastBlock.Hash, data, nonce, difficulty)
		substring := currentHash[0:difficulty]
		if cmp.Equal(strings.Repeat("0", difficulty), substring) {
			break
		}

	}
	block := Block{currentTime, lastBlock.Hash, currentHash, data, nonce, difficulty}
	return block
}
func AdjustDifficulty(block Block, s string, mineRate int) int {

	difficulty := block.Difficulty

	blockTimeStamp, _ := strconv.Atoi(block.Timestamp)
	currentTimeStamp, _ := strconv.Atoi(s)
	i := blockTimeStamp + mineRate

	if i > currentTimeStamp {
		return difficulty + 1
	} else {
		if difficulty != 0 {
			return difficulty - 1
		}
		return difficulty

	}

}

func getHash(timeStamp string, lastHash string, data interface{}, nonce int, difficulty int) string {

	var buffer strings.Builder
	buffer.WriteString(timeStamp)
	buffer.WriteString(lastHash)
	buffer.WriteString(fmt.Sprint(data))
	buffer.WriteString(strconv.Itoa(nonce))
	buffer.WriteString(strconv.Itoa(difficulty))
	return util.GetHash(buffer.String())

}

func GetBlockHash(block Block) string {
	return getHash(block.Timestamp, block.LastHash, block.Data, block.Nonce, block.Difficulty)
}
