package main

import (
	"os/exec"

	"github.com/peterzky/misc/tts/say"
)

func main() {

	xsel := exec.Command("xsel", "-o")
	text, _ := xsel.Output()
	say.Say(string(text))

}
