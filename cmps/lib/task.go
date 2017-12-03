package lib

import (
	"fmt"
	"os"
	"os/exec"
)

type Task interface {
	Start() error
	Stop() error
	Status()
}

type Mpv struct {
	Url string
	cmd *exec.Cmd
}

func NewMpv(url string, proxy bool) *Mpv {
	if proxy {
		mpv := exec.Command("mpv", "--ytdl-format", "mp4", url)
		env := os.Environ()
		env = append(env, fmt.Sprintf("https_proxy=http://localhost:8123"))
		env = append(env, fmt.Sprintf("http_proxy=http://localhost:8123"))
		mpv.Env = env
		return &Mpv{url, mpv}
	} else {
		mpv := exec.Command("mpv", "--ytdl-format", "mp4", url)
		return &Mpv{url, mpv}

	}

}

func (m *Mpv) Stop() error {
	return m.cmd.Process.Kill()
}

func (m *Mpv) Start() error {
	return m.cmd.Start()
}

func (m *Mpv) Status() {
}
