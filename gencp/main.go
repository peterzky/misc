package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

const CFlags string = "NIX_CFLAGS_COMPILE"
const CxxStdLib string = "NIX_CXXSTDLIB_COMPILE"

func main() {
	flags := os.Getenv(CFlags) + os.Getenv(CxxStdLib)

	cmake := exec.Command("cmake", "-DCMAKE_EXPORT_COMPILE_COMMANDS=1")
	cmake.Run()
	fmt.Println("compile_commands.json generated.")
	fmt.Println(".clang_complete generated.")

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
