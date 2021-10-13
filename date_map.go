package main

import (
	"os"
	"path/filepath"
	"time"
	"fmt"
	"github.com/araddon/dateparse"
	"io/ioutil"
)

type DateMap struct {
	FullPath string
	Date     time.Time
}

var getAbs = filepath.Abs
var readDir = os.ReadDir
var parseDate = dateparse.ParseAny

func createDateMapsFromDir(rpath string, tr *timeRange) []*DateMap {
	var result []*DateMap
	abspath, _ := getAbs(rpath)
	dirs, _ := readDir(abspath)
	x := ""

	for _, file := range dirs {
		t, _ := parseDate(file.Name())

		fullPath := filepath.Join(abspath, file.Name())

		if t.After(time.Now().AddDate(tr.months, tr.weeks, tr.days)) {
			content, _ := ioutil.ReadFile(fullPath)
			x += string(content)
		}


		datemap := &DateMap{
			fullPath,
			t,
		}
		result = append(result, datemap)
	}

	fmt.Println(x)
	return result
}
