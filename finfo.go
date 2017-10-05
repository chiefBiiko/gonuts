package main

import (
	"fmt"
	"log"
	"os"
)

const HELP string = "Usage:\tfinfo FILE\n" +
	"FILE\tFile or directory name to display info for.\n"

const TEMPLATE string = "%s: %s\nSize: %s\nModTime: %s\n"

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
		log.Fatalf("error: no argument\n%s", HELP)
	}
	fi, err := os.Stat(os.Args[1])
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	var ftype string
	if fi.IsDir() {
		ftype = "Directory"
	} else {
		ftype = "File"
	}
	fmt.Printf(TEMPLATE,
	  ftype, fi.Name(), prettyBytes(fi.Size()), fi.ModTime().UTC().String())
}
