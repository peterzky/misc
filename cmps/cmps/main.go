package main

import (
	"log"
	"net/http"

	"fmt"

	"github.com/gorilla/mux"
	"github.com/peterzky/misc/cmps/lib"
)

var TaskList []lib.Task

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/video/add/{url}", videoAdd)
	router.HandleFunc("/video/list/", videoList)

	log.Fatal(http.ListenAndServe("127.0.0.1:3111", router))
}

func videoAdd(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url := vars["url"]
	task := lib.NewMpv(url, true)
	task.Start()
	TaskList = append(TaskList, task)
	fmt.Fprintln(w, "playing: ", url)
}

func videoList(w http.ResponseWriter, r *http.Request) {
	var list string
	for key, task := range TaskList {
		switch v := task.(type) {
		case lib.Mpv:
			list += fmt.Sprintf("%d:%s\n", key, v.Url)

		}
	}
	fmt.Fprintln(w, list)
}
