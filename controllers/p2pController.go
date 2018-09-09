package controllers

import (
	"github.com/labstack/echo"

	"bitbucket.org/blockchain/p2pserver"
	"bitbucket.org/blockchain/services"
	"bitbucket.org/blockchain/transaction"
	"bitbucket.org/blockchain/util"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/url"
)

type Message struct {
	KeyV string `json:"key"`
}

type ClientType struct {
	Upgrader websocket.Upgrader
}
type ServerType struct {
}

type ConnectionRequest struct {
	Path   string `json:"path"`
	Port   string `json:"port"`
	Domain string `json:"domain"`
}

var Clients = make(map[*websocket.Conn]bool)
var disrupt = make(chan *websocket.Conn)

func (c *Controller) CreateWebSocketConnection(ctx echo.Context) error {

	clientType := new(ClientType)
	clientType.OpenConnection(ctx.Response(), ctx.Request())
	return nil

}

func (c *Controller) RegisterWebSocket(ctx echo.Context) error {

	serverType := ServerType{}
	request := new(ConnectionRequest)
	ctx.Bind(&request)
	go serverType.OpenConnection(request)
	return ctx.JSON(http.StatusOK, GetSuccessResponse("connected"))

}

func (c *ServerType) OpenConnection(request *ConnectionRequest) {
	u := url.URL{Scheme: "ws", Host: request.Domain + ":" + request.Port, Path: request.Path}
	log.Infoln("creating connection with server with url ", u)
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Error("connection failed with error :", err)
	}
	log.Infoln("connection successfully established")
	Clients[conn] = true
	defer conn.Close()

	go ReadMessages(conn, disrupt)
	go WriteMessages(disrupt)
	handleClientError(disrupt)

}

func handleClientError(disrupt chan *websocket.Conn) {
	connection := <-disrupt
	if connection != nil {
		connection.Close()
		delete(Clients, connection)
	}
}

func (c *ClientType) OpenConnection(w http.ResponseWriter, r *http.Request) {

	c.Upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, e := c.Upgrader.Upgrade(w, r, nil)
	if e != nil {
		log.Error("connection failed with error :", e)
	}
	Clients[conn] = true
	defer conn.Close()
	disrupt := make(chan *websocket.Conn)
	go ReadMessages(conn, disrupt)
	go WriteMessages(disrupt)
	handleClientError(disrupt)

}

func ReadMessages(conn *websocket.Conn, connection chan *websocket.Conn) {
	for {
		data := p2pserver.P2pMessage{}
		err := conn.ReadJSON(&data)
		log.Infoln("reading data from client")
		if err == nil {
			switch data.Type {
			case util.Blocks:
				services.Replace(data.Chain)
				break
			case util.Transactions:
				transaction.UpdateOrAddTransaction(&data.Transaction)
				break
			}

		} else {
			connection <- conn
			log.Error("error while reading the message from client ", err)
		}
	}

}

func WriteMessages(connection chan *websocket.Conn) {
	for {
		msg := <-p2pserver.Broadcast
		if msg != nil {
			for client := range Clients {
				log.Infoln("sending data to client")
				err := client.WriteJSON(msg)
				if err != nil {
					connection <- client
				}
				log.Infoln("data sent")
			}
		} else {
			log.Infoln("nothing to consume")
		}
	}
}

//
//func (c *ClientType) ReadMessage(conn *websocket.Conn, er chan error) {
//
//	for {
//		readMessageAndReplace(conn, er)
//	}
//
//}

//func readMessageAndReplace(conn *websocket.Conn, er chan error) {
//	data := p2pserver.P2pMessage{}
//	err := conn.ReadJSON(&data)
//	log.Infoln("reading data from client")
//	if err == nil {
//
//		switch data.Type {
//
//		case util.Blocks:
//			services.Replace(data.Chain)
//			break
//		case util.Transactions:
//			transaction.UpdateOrAddTransaction(&data.Transaction)
//			break
//		}
//
//	} else {
//		log.Error("error while reading the message from client ", err)
//		er <- err
//	}
//}

//func (c *ClientType) WriteMessage(er chan error) {
//
//	for {
//		msg := <-p2pserver.Broadcast
//		if msg != nil {
//			for client := range Clients {
//				log.Infoln("sending data to client")
//				err := client.WriteJSON(msg)
//				if err != nil {
//					log.Error("error while writing message to client ", err)
//					client.Close()
//					delete(Clients, client)
//				}
//				log.Infoln("data sent")
//			}
//		}
//
//	}
//}

//func (c *ServerType) WriteMessage(conn *websocket.Conn, er chan error) {
//
//	for {
//		msg := <-p2pserver.Broadcast
//		if msg != nil {
//			log.Infoln("writing data to server")
//			err := conn.WriteJSON(msg)
//			if err != nil {
//				log.Info("error while sending data to server ", err)
//				conn.Close()
//				er <- err
//			} else {
//				log.Info("data send to the server")
//			}
//		}
//
//	}
//
//}

//func (c *ServerType) ReadMessage(conn *websocket.Conn, er chan error) {
//	defer conn.Close()
//	for {
//		readMessageAndReplace(conn, er)
//	}
//
//}
