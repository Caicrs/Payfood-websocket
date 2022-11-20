package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func echo(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()
	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Fatal(err)
			break
		}
		fmt.Printf("%s", message)
		err = conn.WriteMessage(mt, message)
		if err != nil {
			log.Fatal(err)
			break
		}
	}
}
func connect(w http.ResponseWriter, r *http.Request) {
	message := "Connected"
	fmt.Printf("recv: %s", message)
}

func main() {
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/", connect)
	http.ListenAndServe(":8000", nil)
}
