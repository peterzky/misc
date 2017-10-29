package main

import (
	"os/exec"

	"sync"

	"fmt"

	"github.com/peterzky/misc/tts/say"
)

func main() {

	xsel := exec.Command("xsel", "-b")
	text, _ := xsel.Output()
	voiceParts := say.Split(string(text), 200)
	for _, vp := range voiceParts {
		fmt.Printf("%d:\n%s\n", vp.Index, vp.Message)
		fmt.Println("----------------")
	}
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
