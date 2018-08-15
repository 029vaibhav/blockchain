package block

import (
	"bitbucket.org/blockchain/time"
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

type Block struct {
	Timestamp string `json:"timestamp"`
	LastHash  string `json:"last_hash"`
	Hash      string `json:"hash"`
	Data      string `json:"data"`
}

func Genesis() Block {
	block := Block{time.MillisInString(), "-----", "firstBlockHash", "first block"}
	return block
}

func MineBlock(lastBlock Block, data string) Block {

	currentTime := time.MillisInString()
	currentHash := getHash(currentTime, lastBlock.Hash, data)
	block := Block{currentTime, lastBlock.Hash, currentHash, data}
	return block
}

func getHash(timeStamp string, lastHash string, data string) string {

	var buffer strings.Builder
	buffer.WriteString(timeStamp)
	buffer.WriteString(lastHash)
	buffer.WriteString(data)
	hash := sha256.New()
	hash.Write([]byte(buffer.String()))
	hashValue := hash.Sum(nil)
	return hex.EncodeToString(hashValue)

}
