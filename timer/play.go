package main

import (
	"os/exec"
	"path"
	"runtime"
)

func play(f string) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}

	cmd := exec.Command("aplay", path.Dir(filename)+"/sound/"+f+".wav")
	cmd.Run()
}
