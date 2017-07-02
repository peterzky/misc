package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func play(addr string, proxy bool) {
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

func route(addr string, sites []string) bool {
	proxy := false
	for _, v := range sites {
		if strings.Contains(addr, v) {
			proxy = true
		}
	}
	return proxy

}

func main() {
	sites := Config("/.gompvrc")
	if len(os.Args) != 2 {
		fmt.Println("Need video address.")
	} else {
		address := os.Args[1]
		proxy := route(address, sites)
		fmt.Println(proxy)
		play(address, proxy)
	}
}
