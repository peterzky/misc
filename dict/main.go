package main

import (
	"fmt"
	"os/exec"

	"github.com/peterzky/misc/dict/lib"
)

const (
	APPID     = "2f871f8481e49b4c"
	APPSECRET = "CQFItxl9hPXuQuVcQa5F2iPmZSbN0hYS"
)

func main() {
	c := &lib.Client{
		AppID:     APPID,
		AppSecret: APPSECRET,
	}

	xsel := exec.Command("xsel", "-o")
	out, err := xsel.Output()
	text := string(out)
	if err != nil {
		panic(err)
	}
	if text == "" {
		panic("no input")
	}

	r, _ := c.Query(text)
	str := r.Format()
	fmt.Println(str)

	lib.DzenAtCursor(str)
}
