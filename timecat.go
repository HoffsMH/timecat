package timecat

import (
	"errors"
	"io/ioutil"
	"path/filepath"
	"time"
	"os"
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
var readDir = os.ReadDir
var parseDate = dateparse.ParseAny
var readFile = func(filename string) (string, error) {
	bytes, err := os.ReadFile(filename)
	return string(bytes), err
}
var now = time.Now
var nowSimpleDate = func() string {
	return now().Format("2006-01-02")
}

var nowISODate = func() string {
	return now().Format(time.RFC3339)
}

func Cat(rpath string, tr *TimeRange) string {
	abspath, _ := getAbs(rpath)
	dirs, _ := readDir(abspath)
	text := ""

	for _, file := range dirs {
		t, err := parseDateFileName(file.Name())
		if err != nil {
			continue
		}
		fullPath := filepath.Join(abspath, file.Name())

		if t.After(time.Now().AddDate(-tr.Months, -tr.Weeks, -tr.Days)) {
			content, _ := ioutil.ReadFile(fullPath)
			text += plainTextHeading + " " + file.Name() + "\n"
			text += string(content) + "\n"
		}
	}
	text += plainTextHeading + " cap.md"

	return text
}

func TimestampString(str string) string {
	_, err := parseDateFileName(str)

	if err != nil {
		return prependCurrentSimpleDate(str)
	}
	return str
}

func prependCurrentSimpleDate(str string) string {
	return nowSimpleDate() + "-" + str
}

func prependCurrentISODate(str string) string {
	return nowISODate() + "-" + str
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
