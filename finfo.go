package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func prettyBytes(numbytes int64) string {
	if numbytes < 1000 {                      // bytes
		return fmt.Sprintf("%d bytes", numbytes)
	} else if numbytes >= 1000 && numbytes <= 1000000 { // KB
		return fmt.Sprintf("%.3f KB", float64(numbytes)/float64(1000))
	} else if numbytes >= 1000000 && numbytes <= 1000000000  { // MB
		return fmt.Sprintf("%.3f MB", float64(numbytes)/float64(1000000))
	} else {                                                          // GB
		return fmt.Sprintf("%.3f GB", float64(numbytes)/float64(1000000000))
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Fprint(os.Stderr, "error: no argument\nUsage:\tfinfo FILE\n" +
			           "FILE\tFile or directory name to display info for.\n")
	  os.Exit(1)
	}
	fi, err := os.Stat(os.Args[1])
	if err != nil {
		panic(err)
	}
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	var fpath string = filepath.Join(cwd, os.Args[1])
	var ftype string
	if fi.IsDir() {
		ftype = "Directory"
	} else {
		ftype = "File"
	}
	fmt.Printf("%s: %s\nSize: %s\nModTime: %s\n",
	  ftype, fpath, prettyBytes(fi.Size()), fi.ModTime().UTC().String())
}
