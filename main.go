package main

import (
	"os"
)

type timeRange struct {
	months int
	weeks int
	days int
}

func main() {
	root := "./testcapdir"
	monthArg := 0
	weekArg := 0
	dayArg := 0

	if len(os.Args) > 1 {
		root = os.Args[1]
	}
	if len(os.Args) > 2 {
		monthArg = os.Args[2]
	}

	if len(os.Args) > 3 {
		weekArg = os.Args[3]
	}

	if len(os.Args) > 4 {
		dayArg = os.Args[4]
	}
	tr := &timeRange{monthArg,weekArg,dayArg}

	createDateMapsFromDir(root, tr)
}

// a function that takes a list of filenames and maps each filename into a date
//
func createDateMapFromDir() {}
