package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	root := "/home/hoffs/somedir"

	fmt.Println("hello from matchcat");

	direntries, _ := os.ReadDir(root);

	for _, direntry := range direntries {
		fmt.Println(direntry.Name());
	}
	x := time.Now().UTC().Format("2006-01-02T15:04:05-0700");
	fmt.Println(x);
	os.Create(x + ".md");
}

func Ok() string {
	return "hi";
}

