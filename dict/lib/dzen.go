package lib

import (
	"io"
	"os/exec"
	"strings"
)

func Dzen(str, x, y, time string) {
	echo := exec.Command("echo", str)
	dzen := exec.Command("dzen2", "-h", "200", "-w", "200", "-x", x, "-y", y, "-p", time,
		"-fn", "WenQuanYi Micro Hei", "-e", "button1=exit", "-l", "4")
	r, w := io.Pipe()
	echo.Stdout = w
	dzen.Stdin = r
	echo.Start()
	dzen.Start()
	echo.Wait()
	w.Close()
	dzen.Wait()
}

func Cursor() (string, string) {
	xdotool := exec.Command("xdotool", "getmouselocation")
	buf, err := xdotool.Output()
	xdotool.Start()
	xdotool.Wait()
	if err != nil {
		panic(err)
	}
	str := string(buf)
	xstr, ystr := strings.Split(str, " ")[0], strings.Split(str, " ")[1]
	x, y := strings.SplitAfter(xstr, ":")[1], strings.SplitAfter(ystr, ":")[1]
	return x, y

}

func DzenAtCursor(msg string) {
	x, y := Cursor()
	Dzen(msg, x, y, "5")
}
