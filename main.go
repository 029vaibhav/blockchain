package main

import (
	"bitbucket.org/blockchain/blockchain"
	"encoding/json"
	"fmt"
)

func main() {

	blockchain.AddBlock("random")
	value := blockchain.GetBlockChain()
	fmt.Println(json.Marshal(value))
	blockchain.AddBlock("random2")
	value = blockchain.GetBlockChain()
	fmt.Println(json.Marshal(value))

}
