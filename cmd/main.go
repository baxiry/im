package main

import (
	"fmt"
	"net/http"

	"github.com/bashery/im"
	"github.com/gorilla/websocket"
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	//if r.Header.Get("Origin")!="http://"+r.Host {http.Error(w,"Origin not allowed",-1);return}

	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
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
