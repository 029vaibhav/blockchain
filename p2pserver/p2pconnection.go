package p2pserver

// connected clients
var Broadcast = make(chan interface{})

//func SendWebSocketRequest(path string) ([]byte, error) {
//
//	u := url.URL{Scheme: "ws", Host: *addr, Path: "/echo"}
//	req, err := http.NewRequest("GET", path, nil)
//	if err != nil {
//		log.Errorln(err)
//		return nil, err
//	}
//	req.Header.Add("Connection", "Upgrade")
//	req.Header.Add("Sec-WebSocket-Extensions", "permessage-deflate; client_max_window_bits")
//	req.Header.Add("Sec-WebSocket-Key", uuid.NewRandom().String())
//	req.Header.Add("Sec-WebSocket-Version", "13")
//	req.Header.Add("Upgrade", "websocket")
//	client := &http.Client{}
//	resp, err := client.Do(req)
//	if err != nil {
//		log.Errorln(err)
//		return nil, err
//	}
//	defer resp.Body.Close()
//	respbody, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Errorln(err)
//		return nil, err
//	}
//	log.Debug("Response Body:", string(respbody))
//	return respbody, nil
//
//}
