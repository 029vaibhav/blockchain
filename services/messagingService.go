package services

import (
	"bitbucket.org/blockchain/p2pserver"
)

func SendMessage(data interface{}) {

	p2pserver.Broadcast <- data

}

func WriteMessage() {

	<-p2pserver.Broadcast
}
