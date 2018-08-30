package block

import (
	"bitbucket.org/blockchain/environment"
	"bitbucket.org/blockchain/time"
	"crypto/sha256"
	"encoding/hex"
	"github.com/google/go-cmp/cmp"
	"strconv"
	"strings"
)

type Block struct {
	Timestamp  string `json:"timestamp"`
	LastHash   string `json:"last_hash"`
	Hash       string `json:"hash"`
	Data       string `json:"data"`
	Nonce      int    `json:"nonce"`
	Difficulty int    `json:"difficulty"`
}

func Genesis() Block {
	block := Block{"Genesis TimeStamp", "-----", "GenesisHash", "Genesis data", 0, environment.Instance().Get("difficulty").(int)}
	return block
}

func MineBlock(lastBlock Block, data string) Block {

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
		return difficulty - 1
	}

}

func getHash(timeStamp string, lastHash string, data string, nonce int, difficulty int) string {

	var buffer strings.Builder
	buffer.WriteString(timeStamp)
	buffer.WriteString(lastHash)
	buffer.WriteString(data)
	buffer.WriteString(strconv.Itoa(nonce))
	buffer.WriteString(strconv.Itoa(difficulty))
	hash := sha256.New()
	hash.Write([]byte(buffer.String()))
	hashValue := hash.Sum(nil)
	return hex.EncodeToString(hashValue)

}

func GetBlockHash(block Block) string {
	return getHash(block.Timestamp, block.LastHash, block.Data, block.Nonce, block.Difficulty)
}
