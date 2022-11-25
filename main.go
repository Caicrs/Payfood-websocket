package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader *websocket.Upgrader = &websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(*http.Request) bool { return true },
}

func ws(w http.ResponseWriter, r *http.Request, hub *Hub, room string) {
	// upgrade http request to websocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	// upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	if err != nil {
		log.Println("Error upgrading request to websocket connection: ", err)
	}

	// new client
	client := &Client{
		Conn: conn,
		Send: make(chan *WSMessage),
		Hub:  hub,
		Room: room,
	}

	// register client to hub
	client.Hub.Register <- client

	go client.Read()
	go client.Write()
}

func main() {
	port := os.Getenv("PORT")
	if strings.Trim(port, " ") == "" {
		port = "8000"
	}

	fs := http.FileServer(http.Dir("./build"))

	http.Handle("/build/", http.StripPrefix("/build/", fs))

	hub := &Hub{
		Clients:    make(map[string]map[*Client]bool),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *WSMessage),
		mutex:      &sync.RWMutex{},
	}
	go hub.Run()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.Redirect(w, r, "/general", http.StatusTemporaryRedirect)
			return
		}

		tmp, err := template.ParseFiles("index.html")
		if err != nil {
			log.Println(err)
			w.Write([]byte("Server Internal Error"))
			return
		}

		channel := strings.Split(r.URL.Path, "/")[1]
		users := len(hub.Clients[channel])

		if err := tmp.Execute(w, map[string]interface{}{
			"channel": strings.Split(r.URL.Path, "/")[1],
			"users":   users,
		}); err != nil {
			log.Println(err)
			w.Write([]byte("Server Internal Error"))
			return
		}
	})

	http.HandleFunc("/ws/", func(w http.ResponseWriter, r *http.Request) {
		t := strings.Split(strings.Trim(r.URL.Path, " "), "/")
		channel := "general"
		if len(t) > 2 && t[2] != "" {
			channel = t[2]
		}
		ws(w, r, hub, channel)
	})

	log.Fatalln(http.ListenAndServe(":"+port, nil))
}
