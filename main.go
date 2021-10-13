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

	if len(os.Args) > 1 {
		root = os.Args[1]
	}
	tr := &timeRange{0,0,-3}

	createDateMapsFromDir(root, tr)
}

// a function that takes a list of filenames and maps each filename into a date
//
func createDateMapFromDir() {}
