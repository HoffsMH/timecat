package timecat

import (
	"os"
	"path/filepath"
	"github.com/araddon/dateparse"
	"time"
	"io/ioutil"
	"regexp"
	"strings"
	"fmt"
)

type TimeRange struct {
	Months int
	Weeks int
	Days int
}

type FileContent struct {
	Name string
	Content string
}

var getAbs = filepath.Abs
var readDir = os.ReadDir
var parseDate = dateparse.ParseAny
var readFile = os.ReadFile

func Cat(rpath string, tr *TimeRange) string {
	abspath, _ := getAbs(rpath)
	dirs, _ := readDir(abspath)
	text := ""

	for _, file := range dirs {
		t, _ := parseDate(file.Name())
		fullPath := filepath.Join(abspath, file.Name())

		if t.After(time.Now().AddDate(-tr.Months, -tr.Weeks, -tr.Days)) {
			content, _ := ioutil.ReadFile(fullPath)
			text += string(content)
		}
	}

	return text
}

var heading = "^!#"

func Split(rpath string) []FileContent {
	abspath, _ := getAbs(rpath)
	content, _ := readFile(abspath)
	lines := strings.Split(string(content), "\n")
	var result []FileContent
	r := regexp.MustCompile(heading + " (.*)")

	for _, line := range lines {
		match := r.FindStringSubmatch(line)

		// we have already found the first header and are now collecting lines
		if len(result) > 0 && len(match) == 0 {
			result[0].Content += line + "\n"
		}
		fmt.Println(testFile)

		// we found a header on the current line
		if len(match) > 1 {
			// start a new header
			result = append([]FileContent{FileContent{ Name: match[1], Content: ""}}, result...)
		}
	}

	return result
}
