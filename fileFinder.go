package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

// PATH is file path
const PATH string = "path"

// PATTETRN is file name pattern
const PATTETRN string = "pattern"

// SHOWTREE is file name pattern
const SHOWTREE string = "showTree"

func main() {
	// 	reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter text: ")
	// text, _ := reader.ReadString('\n')
	// fmt.Println(text)

	argsValues := make(map[string]string)

	for _, arg := range os.Args[1:] {
		if strings.Index(arg, "-") >= 0 && strings.Index(arg, "=") > 0 {

			regstr, _ := regexp.Compile(`^-`)
			arg = regstr.ReplaceAllString(arg, "")

			list := strings.Split(arg, "=")
			argsValues[list[0]] = list[1]
		}
	}

	proccessCommand(argsValues)
}

func proccessCommand(argsValues map[string]string) {
	if argsValues[PATH] != "" && argsValues[SHOWTREE] != "" {
		BuildTree(argsValues[PATH])
	}

	if argsValues[PATTETRN] != "" && argsValues[PATH] != "" {
		findByPattern(argsValues[PATH], argsValues[PATTETRN])
	}
}

// BuildTree build os tree
func BuildTree(path string) {
	pathRead(path, 1)
}

func findByPattern(root string, pattern string) {
	regstr, regexpErr := regexp.Compile(pattern)
	if regexpErr != nil {
		panic(regexpErr)
	}

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if regstr.MatchString(info.Name()) {
			fmt.Println(path)
		}

		return nil
	})

	if err != nil {
		panic(err)
	}
}

func pathRead(root string, lvl int) {
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		value := strings.ReplaceAll(path, root+"/", "")
		if root == path || info.Name() != value {
			return nil
		}

		if info.IsDir() {
			fmt.Println(strings.Repeat(" ", (lvl-1)*3) + value)
			pathRead(path, lvl+1)
		} else {
			fmt.Println(strings.Repeat(" ", (lvl-1)*3) + "|-" + info.Name() + " ")
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

}

func testEx() {
	cmd := exec.Command("ls", "-lah")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))
}
