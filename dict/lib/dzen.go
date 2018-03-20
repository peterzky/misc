package lib

import (
	"io"
	"os/exec"
	"strings"
)

func Dzen(str, x, y, width string) {
	echo := exec.Command("echo", str)
	dzen := exec.Command("dzen2", "-w", width, "-x", x, "-y", y, "-p", "6",
		"-fn", "Sarasa UI SC", "-l", "6",
		"-e", "onstart=uncollapse,scrollhome;"+
			"button1=exit;"+
			"button3=exit;"+
			"button5=scrolldown:1;"+
			"button4=scrollup:1")
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

func DzenAtCursor(msg, width string) {
	x, y := Cursor()
	Dzen(msg, x, y, width)
}

func DoubleClick() {
	action := exec.Command("xdotool", "click", "--repeat", "2", "1")
	action.Run()
}

func MidleClick() {
	action := exec.Command("xdotool", "click", "2")
	action.Run()
}
