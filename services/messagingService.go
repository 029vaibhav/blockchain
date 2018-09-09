package services

import (
	"bitbucket.org/blockchain/p2pserver"
	log "github.com/sirupsen/logrus"
)

func SendMessage(data interface{}) {

	log.Infoln("about to broadcast")
	p2pserver.Broadcast <- data

}

func WriteMessage() {

	<-p2pserver.Broadcast
}
