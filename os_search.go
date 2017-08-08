package os_search

import (
	"fmt"
	"os"
	"path/filepath"
)

const HELP string = "os_search search [start]\n" +
	"search\tFile or directory name to search for.\n" +
	"start\tDirectory name where to start searching.\n\n" +
	"argument search can either be a string or match pattern. go docs:\n" +
	"https://golang.org/pkg/path/filepath/#Match\n"

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

func VisitEntry(epath string, fi os.FileInfo, err error) error {
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
		apath, err := filepath.Abs(epath)
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println(apath)
		os.Exit(0)
	}
	return nil
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("error: no arguments\n")
		fmt.Println(HELP)
		os.Exit(1)
	}
	SEARCH = os.Args[1]
	if len(os.Args) == 2 {
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println("error: os.Getwd() failed")
			os.Exit(1)
		}
		START = cwd
	} else {
		START = os.Args[2]
	}
	exs, err := exists(START)
	if err != nil {
		fmt.Println("error: operating system error")
		os.Exit(1)
	}
	if !exs {
		fmt.Println("error: start directory does not exist\n")
		fmt.Println(HELP)
		os.Exit(1)
	}
	filepath.Walk(START, VisitEntry)
}
