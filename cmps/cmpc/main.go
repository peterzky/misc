package main

import (
	"os"

	"encoding/json"
	"io"
	"net/http"

	"fmt"
	"io/ioutil"

	"github.com/peterzky/misc/cmps/lib"
)

func main() {
	url := os.Args[1]
	data := lib.MpvAdd{url, "play"}
	r, w := io.Pipe()
	json.NewEncoder(w).Encode(data)
	req, err := http.NewRequest("POST", "http://127.0.0.1:3111/video/add/", r)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}
