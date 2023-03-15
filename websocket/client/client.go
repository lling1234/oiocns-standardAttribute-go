package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

func main() {
	// Set up the WebSocket connection
	u := url.URL{Scheme: "ws", Host: "anyinone.com:800", Path: "/orginone/kernel/hub?id=423443007350640640"}
    headers := http.Header{
        "Accept-Encoding": []string{"gzip, deflate"},
        // "Connection": []string{"Upgrade"},
        "Host": []string{"anyinone.com:800"},
        "Origin": []string{"http://anyinone.com:800"},
        "Sec-WebSocket-Key": []string{"XoCn2gD5ok6wXmvh79ss3Q=="},
        "Sec-WebSocket-Extensions": []string{"permessage-deflate; client_max_window_bits"},
        "Sec-WebSocket-Version": []string{"13"},
        // "Upgrade": []string{"websocket"},
    }
    
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), headers)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conn.Close()

	// Start a goroutine to read messages from the WebSocket connection
	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			fmt.Printf("Received message: %s\n", string(message))
		}
	}()

	// Send messages to the WebSocket connection
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	for {
		select {
		case <-interrupt:
			fmt.Println("Interrupt received, exiting...")
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			}
			return
		default:
			message := []byte("Hello, server!")
			err := conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Println("write:", err)
				return
			}
			fmt.Printf("Sent message: %s\n", string(message))
		}
	}
}
