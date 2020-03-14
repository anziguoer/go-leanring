package main

import (
	"fmt"
	"websocket/pkg/websocket"
	"net/http"
)

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Websocket Endpoint Hit")
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}
	pool.Register<-client
	client.Read()
}

func setupRoutes(){
	pool := websocket.NewPool()
	go pool.Start()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}

func main() {
	fmt.Println("Chat App v0.01")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}