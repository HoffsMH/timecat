package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/araddon/dateparse"
)

type DateMap struct {
	FileInfo fs.FileInfo
	FullPath string
	Date     time.Time
}

var getAbs = filepath.Abs
var readDir = os.ReadDir
var parseDate = dateparse.ParseAny

func someFunc() string {
	return "asuh"
}

func createDateMapsFromDir(rpath string) []*DateMap {
	var result []*DateMap
	abspath, _ := getAbs(rpath)
	dirs, _ := readDir(abspath)

	for _, file := range dirs {
		info, _ := file.Info()
		t, _ := parseDate(file.Name())

		datemap := &DateMap{
			info,
			filepath.Join(abspath, file.Name()),
			t,
		}
		result = append(result, datemap)
	}

	return result
}
