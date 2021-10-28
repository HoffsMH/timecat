package timecat

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"github.com/araddon/dateparse"
)

type TimeRange struct {
	Months int
	Weeks  int
	Days   int
}

type FileContent struct {
	Name    string
	Content string
	Dir     string
}

var getAbs = filepath.Abs
var readDir = func(dirname string) []string {
	dirs, _ := os.ReadDir(dirname)
	var result []string

	for _, file := range dirs {
		result = append(result, file.Name())
	}

	return result
}
var parseDate = dateparse.ParseAny
var readFile = func(filename string) (string, error) {
	bytes, err := os.ReadFile(filename)
	return string(bytes), err
}
var now = time.Now

func Cat(rpath string, tr *TimeRange) string {
	abspath, _ := getAbs(rpath)
	files := readDir(abspath)
	text := ""

	for _, file := range files {
		t, err := parseDateFileName(file)
		if err != nil {
			continue
		}
		fullPath := filepath.Join(abspath, file)

		if t.After(now().AddDate(-tr.Months, -tr.Weeks, -tr.Days)) {
			content, _ := readFile(fullPath)
			text += plainTextHeading + " " + file + "\n"
			text += ensureNewline(string(content))
		}
	}
	text += plainTextHeading + " cap.md\n"

	return text
}

func ensureNewline(s string) string {
	match, _ := regexp.MatchString("\n$", s)
	if match == true {
		return s
	}
	return s + "\n"
}

func filterFiles(filenames []string, searchtext string) []string {
	var result []string
	for _, f := range filenames {
		match, _ := regexp.MatchString(searchtext, f)
		if match == true {
			result = append(result, f)
		}
	}
	return result
}

func parseDateFileName(fn string) (time.Time, error) {
	if len(fn) < 10 {
		return now(), errors.New("not long enough to contain a date")
	}

	datePortion := fn[:10]
	dateOutput, err := parseDate(datePortion)
	if err == nil {
		return dateOutput, nil
	}

	if len(fn) < 24 {
		return now(), errors.New("No date detected")
	}

	datePortion = fn[:25]
	dateOutput, err = parseDate(datePortion)
	if err == nil {
		return dateOutput, nil
	}
	return now(), errors.New("No date detected")
}
