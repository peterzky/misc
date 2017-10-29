package main

import (
	"fmt"
	"os/exec"

	"sync"

	"flag"

	"github.com/peterzky/misc/tts/say"
)

var clip, selec, debug bool
var input string

func init() {
	flag.BoolVar(&clip, "clip", false, "send clipboard")
	flag.BoolVar(&selec, "sel", false, "send selection")
	flag.BoolVar(&debug, "debug", false, "debug output")
	flag.StringVar(&input, "t", "", "send text")
}

func main() {
	flag.Parse()
	var text string

	if clip {
		xsel := exec.Command("xsel", "-o")
		out, err := xsel.Output()
		text = string(out)
		if err != nil {
			panic(err)
		}
	}

	if selec {
		xsel := exec.Command("xsel", "-o")
		out, err := xsel.Output()
		text = string(out)
		if err != nil {
			panic(err)
		}

	}
	if !clip && !selec && input != "" {
		text = input
	}

	voiceParts := say.Split(text, 150)
	if debug {
		for _, v := range voiceParts {
			fmt.Printf("Index: %d\nMessage: %s\nFileName: %s\n", v.Index, v.Message, v.FileName)
			fmt.Println("---------------------------")
		}
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
