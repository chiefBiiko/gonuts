package main

import (
	"os"
	"path/filepath"
)

var HELP_FLAGS = []string{"h", "-h", "--h", "help", "-help", "--help"}

const HELP string = "\nUsage:\tphind SEARCH [START]\n" +
	"SEARCH\tFile or directory to search for.\n" +
	"START\tDirectory where to start searching; " +
	"default: current working directory.\n\n" +
	"Argument SEARCH can either be a string or match pattern. go docs:\n" +
	"https://golang.org/pkg/path/filepath/#Match.\n"

func contains(slize []string, item string) bool {
	for _, v := range slize {
		if v == item {
			return true
		}
	}
	return false
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || !os.IsNotExist(err)
}

func visitEntry(epath string, fi os.FileInfo, err error) error {
	if err != nil {
		os.Stderr.WriteString("error: " + err.Error() + "\n")
		return nil
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

var SEARCH, START string

func main() {
	switch len(os.Args) {
	case 1:
		os.Stderr.WriteString("error: no arguments\n" + HELP)
		os.Exit(1)
	case 2:
		if contains(HELP_FLAGS, os.Args[1]) {
			os.Stdout.WriteString(HELP)
			os.Exit(0)
		}
		START, _ = os.Getwd()
	default:
		START = os.Args[2]
	}
	exs := exists(START)
	if !exs {
		os.Stderr.WriteString("error: START does not exist\n" + HELP)
		os.Exit(1)
	}
	SEARCH = os.Args[1]
	filepath.Walk(START, visitEntry)
	os.Exit(0)
}
