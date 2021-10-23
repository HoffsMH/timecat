package timecat

import (
	"os"
	"path"
	"regexp"
	"strings"
)

var writeFile = os.WriteFile

var plainTextHeading = "##"
var heading = "^" + plainTextHeading

func Split(rpath string, dir string) []FileContent {
	fileAbsPath, _ := getAbs(rpath)
	dirAbsPath, _ := getAbs(dir)
	content, _ := readFile(fileAbsPath)
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
			result = append([]FileContent{FileContent{Dir: dirAbsPath, Name: match[1], Content: ""}}, result...)
		}
	}

	return result
}

func WriteSplits(fcs []FileContent) {
	for _, fc := range fcs {
		fileName := fc.Name
		if _, err := parseDateFileName(fc.Name); err != nil {
			fileName = prependCurrentISODate(fc.Name)
		}

		writeFile(path.Join(fc.Dir, fileName), []byte(fc.Content), 0644)
	}
}
