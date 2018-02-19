package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func genconf(fname string) {
	var str string
	str = includeStr("gcc")
	if str == "" {
		str = includeStr("clang")
	}
	var fitterdList []string

	for _, flag := range strings.Split(str, " ") {
		matched, err := regexp.MatchString("^/nix.*", flag)
		if err != nil {
			panic(err)
		}

		if matched {
			fitterdList = append(fitterdList, flag)
		}
	}

	file, err := os.Create(".cquery")

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

func includeStr(compiler string) string {
	cc := exec.Command(compiler, "-v", "-E", "-")
	out, _ := cc.CombinedOutput()
	split := strings.SplitAfter(string(out), "#include <...>")
	if len(split) >= 2 {
		return split[1]
	} else {
		return ""
	}
}

func main() {
	cmake := exec.Command("cmake", "-DCMAKE_EXPORT_COMPILE_COMMANDS=1")
	cmake.Run()
	fmt.Println("compile_commands.json generated.")
	fmt.Println(".clang_complete generated.")
	genconf(".cquery")
}
