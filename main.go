package main

import (
	"fmt"
	"os"
)

func main() {
	root := "./testcapdir"

	if len(os.Args) > 1 {
		root = os.Args[1]
	}

	datemaps := createDateMapsFromDir(root)

	for _, datemap := range datemaps {
		info := datemap.FileInfo
		fmt.Println(datemap)
		fmt.Println(*datemap)
		fmt.Println(info)
		fmt.Println(info.Name())
	}

	// for _, direntry := range direntries {
	// 	fullPath := filepath.Join(path, direntry.Name())
	// 	stat, _ := os.Stat(fullPath)

	// 	fmt.Println(stat.ModTime())
	// 	fmt.Println("Here is fullpath")
	// 	fmt.Println(fullPath)

	// 	fmt.Println("Here is basename")
	// 	fmt.Println(direntry.Name())
	// 	t, _ := dateparse.ParseAny(direntry.Name())
	// 	fmt.Println(t)
	// }
	// x := time.Now().UTC().Format("2006-02-02T15:04:05-0700")
	// fmt.Println(x)
	// os.Create(x + ".md")
}

// a function that takes a list of filenames and maps each filename into a date
//
func createDateMapFromDir() {}
