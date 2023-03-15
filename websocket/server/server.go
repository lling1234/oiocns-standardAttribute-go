package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func echo(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to a WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// Read and write messages to the WebSocket connection
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Received message: %s\n", string(p))

		err = conn.WriteMessage(messageType, p)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Sent message: %s\n", string(p))
	}
}

func main() {
	http.HandleFunc("/", echo)
	fmt.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", nil)
}
