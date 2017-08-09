package main

import (
	"os"
	"path/filepath"
)

const HELP string = "\nUsage:\tphind SEARCH [START]\n" +
	"SEARCH\tFile or directory to search for.\n" +
	"START\tDirectory where to start searching; " +
	"default: current working directory.\n\n" +
	"Argument SEARCH can either be a string or match pattern. go docs:\n" +
	"https://golang.org/pkg/path/filepath/#Match.\n"

var SEARCH, START string

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || !os.IsNotExist(err)
}

func visitEntry(epath string, fi os.FileInfo, err error) error {
	if err != nil {
		os.Stderr.WriteString("error: " + err.Error() + "\n") // can't walk here
		return nil                                            // but continue
	}
	matched, err := filepath.Match(SEARCH, fi.Name())
	if err != nil {
		os.Stderr.WriteString("error: " + err.Error() + "\n")
		os.Exit(1)
	}
	if matched {
		apath, err := filepath.Abs(epath)
		if err != nil {
			os.Stderr.WriteString("error: " + err.Error() + "\n")
			os.Exit(1)
		}
		os.Stdout.WriteString(apath + "\n")
		os.Exit(0)
	}
	return nil
}

func main() {
	switch len(os.Args) {
	case 1:
		os.Stderr.WriteString("error: no arguments\n" + HELP)
		os.Exit(1)
	case 2:
		START, _ = os.Getwd()
	default:
		START = os.Args[2]
	}
	exs := exists(START)
	if !exs {
		os.Stderr.WriteString("error: start directory does not exist\n" + HELP)
		os.Exit(1)
	}
	SEARCH = os.Args[1]
	filepath.Walk(START, visitEntry)
	os.Exit(0)
}
