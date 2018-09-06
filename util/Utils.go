package util

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/pborman/uuid"
)

func RandomIdGen() string {
	return uuid.NewRandom().String()
}

func GetHash(data string) string {

	hash := sha256.New()
	hash.Write([]byte(data))
	hashValue := hash.Sum(nil)
	return hex.EncodeToString(hashValue)
}
