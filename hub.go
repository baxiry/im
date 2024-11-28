package im

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/tidwall/gjson"
)

// note: income data must be json as `{"event":"","msg":""}`
// event must be subscribe, unsubscriber, close or msg
func ServeMessages(conn *websocket.Conn) {

	var mt sync.Mutex
	for {

		i, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message no.", i)
			conn.Close()
			continue // return
		}

		// un/subscribe if event == un/subscribe.
		income := gjson.ParseBytes(msg[1:])
		event := income.Get("event").Str
		data := income.Get("data").Raw
		channel := income.Get("channel").Str

		switch event {
		case "message":
			publish(i, channel, []byte(data))

		case "subscribe":
			subscribe(channel, conn)
			msg = append([]byte("successfully subscribed to "), []byte(channel)...)

		case "unsubscribe":
			unsubscribe(channel, conn)
			msg = append([]byte("successfully unsubscribed to "), []byte(channel)...)
		default:
			log.Println("unkown event: ", event)
		}

		mt.Lock()
		if err = conn.WriteMessage(i, msg); err != nil {
			log.Println(err)
			mt.Unlock()
			continue
		}
		mt.Unlock()
	}
}
