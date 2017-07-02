package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strings"
)

func Config(config string) []string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	rc := usr.HomeDir + config

	if _, err := os.Stat(rc); os.IsNotExist(err) {
		os.Create(rc)
	}

	dat, err := ioutil.ReadFile(rc)
	if err != nil {
		panic(err)
	}
	s := strings.Split(string(dat), "\n")
	return s
}
