package timecat

import (
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

var writeFile = os.WriteFile

var plainTextHeading = "##"
var heading = "^" + plainTextHeading

var nowISODate = func() string {
	return now().Format(time.RFC3339)
}

func prependCurrentISODate(str string) string {
	return nowISODate() + "-" + str
}

func Split(rpath string, dir string) []FileContent {
	lines := getLines(rpath)

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
			result = append([]FileContent{newFileContent(rpath, dir)}, result...)
		}
	}
	return pruneEmptyFileContents(result)
}

func pruneEmptyFileContents(fcs []FileContent) []FileContent {
	var pruned []FileContent
	r := regexp.MustCompile(".")

	for _, fc := range fcs {
		if match := r.FindStringSubmatch(fc.Content); len(match) > 0 {
			pruned = append(pruned, fc)
		} else {
			os.Remove(filepath.Join(fc.Dir, fc.Name))
		}
	}

	return pruned
}

func getLines(path string) []string {
	absPath, _ := getAbs(path)
	text, _ := readFile(absPath)
	return strings.Split(text, "\n")
}

func newFileContent(name string, dir string) FileContent {
	dir, _ = getAbs(dir)
	if _, err := parseDateFileName(name); err != nil {
		name = prependCurrentISODate(name)
	}

	return FileContent{
		Dir:     dir,
		Name:    path.Base(name),
		Content: "",
	}
}

func WriteSplits(fcs []FileContent) {
	for _, fc := range fcs {
		writeFile(path.Join(fc.Dir, fc.Name), []byte(fc.Content), 0644)
	}
}
