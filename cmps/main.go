package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var DefaultConfig = []Task{
	Mpv{"/mpv/", "POST"},
	Tmux{"/youtube-dl/", "POST", "youtube-dl", "Download"},
}

func main() {
	router := mux.NewRouter()
	for _, task := range DefaultConfig {
		log.Println(task)
		router.HandleFunc(task.Bind(), task.Handler()).Methods(task.Method())
	}

	log.Fatal(http.ListenAndServe("127.0.0.1:3111", router))
}
