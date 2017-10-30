package main

import (
	"fmt"
	"os/exec"

	"flag"

	"os"
	"strings"

	"time"

	"github.com/peterzky/misc/dict/lib"
)

const (
	APPID     = "2f871f8481e49b4c"
	APPSECRET = "CQFItxl9hPXuQuVcQa5F2iPmZSbN0hYS"
)

var sel bool
var width string

func init() {
	flag.BoolVar(&sel, "s", false, "whither using commandline input")
	flag.StringVar(&width, "w", "300", "width of popup")
}

func main() {
	flag.Parse()
	c := &lib.Client{
		AppID:     APPID,
		AppSecret: APPSECRET,
	}

	if sel {
		lib.DoubleClick()
		time.Sleep(time.Millisecond)
		xsel := exec.Command("xsel", "-o")
		out, err := xsel.Output()
		text := string(out)
		if err != nil {
			panic(err)
		}
		if text == "" {
			panic("no input")
		}

		r, err := c.Query(text)
		lib.MidleClick()

		if err != nil {
			lib.DzenAtCursor(err.Error(), width)
		}
		str := r.Format()
		fmt.Println(str)

		lib.DzenAtCursor(str, width)
	} else {
		text := strings.Join(os.Args[1:], "")
		r, err := c.Query(text)
		if err != nil {
			panic(err)
		}
		fmt.Println(r.Format())

	}
}
