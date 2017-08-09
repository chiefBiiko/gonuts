package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const HELP string = "\nUsage:\tphind SEARCH [START]\n" +
	"SEARCH\tFile or directory name to search for.\n" +
	"START\tDirectory name where to start searching.\n\n" +
	"Argument SEARCH can either be a string or match pattern. go docs:\n" +
	"https://golang.org/pkg/path/filepath/#Match.\n" +
	"Argument START defaults to the current working directory."

var SEARCH, START string

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, err
	}
	return false, err
}

func visitEntry(epath string, fi os.FileInfo, err error) error {
	if err != nil {
		fmt.Println("error:", err) // can't walk here
		return nil                 // but continue walking elsewhere
	}
	//if fi.IsDir() { return nil }
	matched, err := filepath.Match(SEARCH, fi.Name())
	if err != nil {
		fmt.Println("error:", err)
		return err // fatal error, guess execution stops here
	}
	if matched {
		apath, err := filepath.Abs(epath)
		if err != nil {
			fmt.Println("error:", err)
			return err
		}
		fmt.Println(apath)
		os.Exit(0)
	}
	return nil
}

func main() {
	switch len(os.Args) {
	case 1:
		fmt.Println("error: no arguments\n", HELP)
		os.Exit(1)
	case 2:
		START, _ = os.Getwd()
	default:
		START = os.Args[2]
	}
	exs, err := exists(START)
	if !exs || err != nil {
		fmt.Println("error: start directory does not seem to exist\n", HELP)
		os.Exit(1)
	}
	SEARCH = os.Args[1]
	filepath.Walk(START, visitEntry)
}
