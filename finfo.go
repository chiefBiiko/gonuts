package main

import (
	"fmt"
	"os"
)

const HELP string = "\nUsage:\tfinfo FILE\n" +
	"FILE\tFile or directory name to display info for."

const TMPL string = "%s: %s\nSize: %s\nModTime: %v\n"

// TODO: fix!??!!
func prettyBytes(numbytes int64) string {
	if numbytes < 1000 { // bytes
		return fmt.Sprintf("%d bytes", numbytes)
	} else if numbytes >= 1000 { // KB
		return fmt.Sprintf("%v KB", numbytes/1000)
	} else if numbytes >= 1000000  { // MB
		return fmt.Sprintf("%v MB", numbytes/1000000)
	} else {                             // GB
		return fmt.Sprintf("%v GB", numbytes/1000000000)
	}
}

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
	fmt.Printf(TMPL, etype, name, prettyBytes(size), modtime)
}
