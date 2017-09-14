package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

func mpv(addr string, proxy bool) {
	if proxy {
		mpv := exec.Command("mpv", "--ytdl-format", "mp4", addr)
		env := os.Environ()
		env = append(env, fmt.Sprintf("https_proxy=http://localhost:8123"))
		env = append(env, fmt.Sprintf("http_proxy=http://localhost:8123"))
		mpv.Env = env
		mpv.Run()
	} else {
		mpv := exec.Command("mpv", addr)
		mpv.Run()
	}
}

func play(w http.ResponseWriter, r *http.Request) {
	re := strings.NewReplacer("/play/", "", "https:/", "https://", "http:/", "http://")
	uri := re.Replace(r.RequestURI)
	decodedUrl, _ := url.QueryUnescape(uri)
	fmt.Printf("playing %s\n", decodedUrl)
	fmt.Fprintf(w, "playing %s\n", decodedUrl)
	mpv(decodedUrl, true)
}

func main() {
	http.HandleFunc("/", play)
	http.ListenAndServe("127.0.0.1:3111", nil)
}
