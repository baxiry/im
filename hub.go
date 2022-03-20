package im

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/tidwall/gjson"
)

//note: income data must be json as `{"event":"","msg":""}`
// event must be subscribe, unsubscriber, close or msg
func ServeMessages(conn *websocket.Conn) {
	fmt.Println("version 0.0.2")

	var mt sync.Mutex
	for {

		i, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message no.", i)
			conn.Close()
			return
		}

		// un/subscribe if event == un/subscribe.
		var smsg = string(msg)
		event := gjson.Get(smsg, "event").String()
		// TODO continue if no event.
		channel := gjson.Get(smsg, "channel").String()
		data := gjson.Get(smsg, "data").String()

		if event == "message" {

			Publish(i, channel, []byte(data))

		} else if event == "subscribe" {

			Subscribe(channel, conn)
			msg = []byte("subscribe to " + channel + " success!")

		} else if event == "unsubscribe" {

			Unsubscribe(channel, conn)
			msg = []byte("unsubscribe from " + channel + " success!")
		}

		fmt.Println(string(msg))

		mt.Lock()
		if err = conn.WriteMessage(i, msg); err != nil {
			fmt.Println(err)
		}
		mt.Unlock()
	}
}
