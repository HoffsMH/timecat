package timecat

import (
	"testing"
)

var testFile = `
asuh
ok
`

func TestWriteSplits(t *testing.T) {
	var testFile = `asuh
ok
asdf
asdf
!# testfile.md
testtext1
testtext2
testtext3
`

	oldReadFile := mockReadFile(testFile, nil)
	defer func() { readFile = oldReadFile }()

	result := Split("testfile", "testdir")
	WriteSplits(result)
}
