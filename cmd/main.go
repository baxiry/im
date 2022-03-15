package main

import (
	"im"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var mt sync.Mutex

func wsHandler(w http.ResponseWriter, r *http.Request) {
	//if r.Header.Get("Origin")!="http://"+r.Host {http.Error(w,"Origin not allowed",-1);return}

	conn, err := websocket.Upgrade(w, r, w.Header(), 2, 2) //1024, 1024)
	if err != nil {
		http.Error(w, "Could not open websocket connection", 404)
	}
	go im.ServeMessages(conn)
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	panic(http.ListenAndServe(":8080", nil))
}
