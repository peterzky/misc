package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

const CFlags string = "NIX_CFLAGS_COMPILE"
const CxxStdLib string = "NIX_CXXSTDLIB_COMPILE"

func main() {
	flags := os.Getenv(CFlags) + os.Getenv(CxxStdLib)

	var fitterdList []string

	for _, flag := range strings.Split(flags, " ") {
		matched, err := regexp.MatchString("^/nix.*", flag)
		if err != nil {
			panic(err)
		}

		if matched {
			fitterdList = append(fitterdList, flag)
		}
	}

	file, err := os.Create(".clang_complete")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	for _, flag := range fitterdList {
		str := fmt.Sprintf("-I%s\n", flag)
		file.WriteString(str)
		fmt.Printf(str)
	}

}
