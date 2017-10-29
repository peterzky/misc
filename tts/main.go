package main

import (
	"os/exec"

	"sync"

	"github.com/peterzky/misc/tts/say"
)

func main() {

	xsel := exec.Command("xsel", "-o")
	text, _ := xsel.Output()
	voiceParts := say.Split(string(text), 150)
	var wg sync.WaitGroup

	for _, vp := range voiceParts {
		wg.Add(1)
		go func(vp say.VoicePart) {
			defer wg.Done()
			say.Download(vp)
		}(vp)
	}
	wg.Wait()
	for _, vp := range voiceParts {
		say.Play(vp)
	}

}
