package timecat

import (
	"fmt"
	"path"
	"strings"
)

// from the complete outside
// when given nothing we get from clipboard
// when given text we write the text to a timestamped iso file in our directory

// ability to provide sub file name?
func Cap(dir string, content []string) {
	fileName := prependCurrentISODate("cap.md")
	absPath, _ := getAbs(dir)
	textContent := strings.Join(content, " ") + "\n"
	writeFile(path.Join(absPath, fileName), []byte(textContent), 0644)
	fmt.Println(path.Join(absPath, fileName))
}
