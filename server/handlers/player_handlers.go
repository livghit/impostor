package handlers

import "net/http"

func HandleCreatePlayer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("We create a player now"))
}
