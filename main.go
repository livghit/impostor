package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()

	log.Println("New WebSocket connection")

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Connection closed:", err)
			break
		}

		fmt.Printf("Received: %s\n", msg)

		// Echo the message back
		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)

	// Serve Vue frontend
	fs := http.FileServer(http.Dir("./ui/dist"))
	http.Handle("/", fs)

	// Route responsible for creating a player
	http.HandleFunc("POST /create-player", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("We create a player now"))
	})
	// Route responsible for creating a lobby
	http.HandleFunc("POST /create-lobby", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("We create a lobby now"))
	})
	// Route responsible for creating a game
	http.HandleFunc("POST /create-game", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("We create a game now"))
	})

	// Server
	port := "8080"
	log.Println("Server started on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
