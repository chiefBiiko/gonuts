package main

import (
	"fmt"
	"os"
)

const HELP string = "\nUsage:\tfinfo FILE\n" +
	"FILE\tFile or directory name to display info for."

const TMPL string = "%s: %s\nSize: %s\nModTime: %v\n"

// TODO: implementation
/*
func prettyBytes(numbytes uint) string {
	if numbytes <= 1000 { // KB
		return fmt.Sprintf("%d bytes", numbytes)
	} else if numbytes <= 1000000 { // MB
		return fmt.Sprintf("%d MB", numbytes/1000)
	} else { // GB  // if numbytes <= 1000000000
		return fmt.Sprintf("%d GB", numbytes/1000000)
	}
}
*/

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
