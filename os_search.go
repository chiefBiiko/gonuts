package os_search

import (
	"fmt"
	"os"
	"path/filepath"
)

var SEARCH, START string

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, err
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func VisitEntry(fp string, fi os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(err) // can't walk here
		return nil       // but continue walking elsewhere
	}
	//if fi.IsDir() { return nil }
	matched, err := filepath.Match(SEARCH, fi.Name())
	if err != nil {
		fmt.Println(err)
		return err // fatal error, guess execution stops here
	}
	if matched {
		afp, err := filepath.Abs(fp)
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println(afp)
		os.Exit(0)
	}
	return nil
}

const HELP string = "os_search search [start]\n" +
	"search\tFile or directory name to search for.\n" +
	"start\tDirectory name where to start searching.\n\n" +
	"argument search can either be a string or match pattern. go docs:\n" +
	"https://golang.org/pkg/path/filepath/#Match\n"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("error: invalid arguments\n")
		fmt.Println(HELP)
		os.Exit(1)
	} else if len(os.Args) == 2 {
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println("error: os.Getwd() failed\n")
			os.Exit(1)
		}
		START = cwd
	} else {
		START = os.Args[2]
	}
	exs, err := exists(START)
	if err != nil {
		fmt.Println("error: operating system error\n")
		os.Exit(1)
	}
	if !exs {
		fmt.Println("error: start directory does not exist\n")
		fmt.Println(HELP)
		os.Exit(1)
	}
	SEARCH = os.Args[1]
	filepath.Walk(START, VisitEntry)
}
