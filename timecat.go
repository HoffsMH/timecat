package timecat

import (
	"os"
	"path/filepath"
	"github.com/araddon/dateparse"
	"time"
	"io/ioutil"
	"regexp"
	"strings"
	"path"
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
	Dir string
}

var getAbs = filepath.Abs
var readDir = os.ReadDir
var parseDate = dateparse.ParseAny
var readFile = func (filename string) (string, error) {
	bytes, err := os.ReadFile(filename);
	return string(bytes), err;
}
var writeFile = os.WriteFile

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
	lines := strings.Split(content, "\n")
	var result []FileContent
	r := regexp.MustCompile(heading + " (.*)")

	for _, line := range lines {
		match := r.FindStringSubmatch(line)

		// we have already found the first header and are now collecting lines
		if len(result) > 0 && len(match) == 0 {
			result[0].Content += line + "\n"
		}

		// we found a header on the current line
		if len(match) > 1 {
			// start a new header
			result = append([]FileContent{FileContent{ Dir: abspath, Name: match[1], Content: ""}}, result...)
		}
	}

	return result
}

func WriteSplits(fcs []FileContent) {
	for _, fc := range fcs {
		writeFile(path.Join(fc.Dir, fc.Name), []byte(fc.Content), 0644)
	}
}

func TimstampString(n string) string {
	name, err := dateparse.ParseAny(n);
	fmt.Println("HERE IS ERR")
	fmt.Println(err)

	fmt.Println("HERE IS NAME")
	fmt.Println(name)

	fmt.Println("OUR ATTEMPT @ SLICING")
	fmt.Println(string([]byte(n)[:10]))

	if err == nil {
		return n;
	}
	return time.Now().Format("2006-01-02") + "-" + n
}
