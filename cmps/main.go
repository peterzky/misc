package main

import (
	"log"
	"net/http"

	"fmt"

	"github.com/gorilla/mux"
)

const proxy = "http_proxy=http://127.0.0.1:8123 https_proxy=http://127.0.0.1:8123"

var DefaultConfig = []Task{
	&Tmux{"/mpv/", "POST", "mpv", fmt.Sprintf("%s mpv --ytdl-format mp4", proxy)},
	&Tmux{"/youtube-dl/", "POST", "Download", "youtube-dl"},
}

func main() {
	router := mux.NewRouter()
	for _, task := range DefaultConfig {
		log.Println(task)
		router.HandleFunc(task.Bind(), task.Handler()).Methods(task.Method())
	}

	log.Fatal(http.ListenAndServe("127.0.0.1:3111", router))
}
