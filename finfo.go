package main

import (
	"fmt"
	"os"
)

const HELP string = "\nUsage:\tfinfo FILE\n" +
	"FILE\tFile or directory name to display info for."

const TMPL string = "%s: %s\nSize: %d bytes\nModTime: %v\n"

func main() {
	if len(os.Args) == 1 {
		fmt.Println("error: no argument")
		fmt.Println(HELP)
		os.Exit(1)
	}
	fi, err := os.Stat(os.Args[1])
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	name, size, modtime := fi.Name(), fi.Size(), fi.ModTime()
	var etype string
	if fi.IsDir() {
		etype = "Directory"
	} else {
		etype = "File"
	}
	fmt.Printf(TMPL, etype, name, size, modtime)
}
