package main

import (
	"log"
	"net/http"

	"fmt"

	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/peterzky/misc/cmps/lib"
)

var TaskList []lib.Task

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/video/add/", videoAdd).Methods("POST")
	router.HandleFunc("/video/list/", videoList)

	log.Fatal(http.ListenAndServe("127.0.0.1:3111", router))
}

func videoAdd(w http.ResponseWriter, r *http.Request) {
	add := new(lib.MpvAdd)
	if err := json.NewDecoder(r.Body).Decode(add); err != nil {
		log.Fatal(err)
	}
	task := lib.NewMpv(add.Url, true)
	task.Start()
	TaskList = append(TaskList, task)
	fmt.Fprintln(w, "playing: ", task.Url)
}

func videoList(w http.ResponseWriter, r *http.Request) {
	var list string
	for key, task := range TaskList {
		switch v := task.(type) {
		case lib.Mpv:
			list += fmt.Sprintf("%d:%s\n", key, v.Url)

		}
	}
	json.NewEncoder(w).Encode(TaskList)
}
