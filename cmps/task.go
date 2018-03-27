package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

type HttpHandler func(http.ResponseWriter, *http.Request)

type Task interface {
	Bind() string
	Method() string
	Handler() HttpHandler
}

// get url from post body
func GetUrl(r *http.Request) string {
	var req struct{ Url string }
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		log.Fatal("can't parse json: ", err)
	}
	return req.Url
}

// play url with mpv
type Mpv struct {
	bind   string
	method string
}

func (m Mpv) Bind() string { return m.bind }

func (m Mpv) Method() string { return m.method }

func (m Mpv) Handler() HttpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		url := GetUrl(r)
		mpv := exec.Command("mpv", "--ytdl-format", "mp4", url)
		env := os.Environ()
		env = append(env, fmt.Sprintf("https_proxy=http://localhost:8123"))
		env = append(env, fmt.Sprintf("http_proxy=http://localhost:8123"))
		mpv.Env = env
		mpv.Start()
		fmt.Fprintln(w, "playing: ", url)
		log.Printf("[mpv] %s", url)
	}
}

// testing dummy
type Dummy struct {
	bind   string
	method string
}

func (d Dummy) Bind() string { return d.bind }

func (d Dummy) Method() string { return d.method }

func (d Dummy) Handler() HttpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "dummy reply!")
		log.Printf("dummy received\n")
	}
}

// execute tmux command
type Tmux struct {
	bind    string
	method  string
	window  string
	command string
}

func (t Tmux) Bind() string { return t.bind }

func (t Tmux) Method() string { return t.method }

func (t Tmux) Handler() HttpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		url := GetUrl(r)
		cmd := fmt.Sprintf("%s '%s'", t.command, url)
		tmux := exec.Command("tmux", "new-window", "-n", t.window, cmd)
		err := tmux.Start()

		if err != nil {
			log.Println(err)
		}
		fmt.Fprintf(w, "downloading: ", url)
		log.Printf("[%s] %s", t.window, url)

	}
}
