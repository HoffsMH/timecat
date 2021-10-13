package main

import (
	"os"
	"path/filepath"
	"time"
	"fmt"
	"github.com/araddon/dateparse"
	"io/ioutil"
)

var getAbs = filepath.Abs
var readDir = os.ReadDir
var parseDate = dateparse.ParseAny

func main(rpath string, tr *timeRange) string {
	abspath, _ := getAbs(rpath)
	dirs, _ := readDir(abspath)
	text := ""

	for _, file := range dirs {
		t, _ := parseDate(file.Name())
		fullPath := filepath.Join(abspath, file.Name())

		if t.After(time.Now().AddDate(tr.months, tr.weeks, tr.days)) {
			content, _ := ioutil.ReadFile(fullPath)
			text += string(content)
		}
	}

	fmt.Println(text)
	return text
}
