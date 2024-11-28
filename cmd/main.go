package main

import (
	"fmt"
	"net/http"

	"github.com/baxiry/im"
	"github.com/gorilla/websocket"
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	//if r.Header.Get("Origin")!="http://"+r.Host {http.Error(w,"Origin not allowed",-1);return}

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not open websocket connection", 404)
	}
	go im.ServeMessages(conn)
}

func main() {

	fmt.Println("version 0.0.2\nim start at :8080")
	http.HandleFunc("/ws", wsHandler)

	panic(http.ListenAndServe(":8080", nil))

}
