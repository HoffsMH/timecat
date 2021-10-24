package timecat

import (
	"testing"
)

func TestSplitWithEmptyFile(t *testing.T) {
	var testFile = ``

	oldReadFile := mockReadFile(testFile, nil)
	defer func() { readFile = oldReadFile }()

	// when split is run with a completely empty file
	// the resulting array of files to write should be empty
	result := Split("testfile", "testdir")

	if len(result) != 0 {
		t.Fatal("there was more or less than one heading")
	}
}

func TestSplitWithNoHeadings(t *testing.T) {
	var testFile = `asuh
ok
asdf
asdf

`

	oldReadFile := mockReadFile(testFile, nil)
	defer func() { readFile = oldReadFile }()

	// when split is run with a file that doens't have a single heading
	// the resulting array of files to write should be empty
	result := Split("testfile", "testdir")
	if len(result) > 0 {
		t.Fatal("there were headings when there should be none")
	}
}

func TestSplitWithOneHeading(t *testing.T) {
	var testFile = `asuh
ok
asdf
asdf
` + plainTextHeading + ` testfile.md
testtext1
testtext2
testtext3
`

	oldReadFile := mockReadFile(testFile, nil)
	defer func() { readFile = oldReadFile }()

	oldNowISODate := mockNowISODate("test-prefix")
	defer func() { nowISODate = oldNowISODate }()

	result := Split("testfile.md", "testdir")

	// when split has some text before a single heading
	// the resulting array has a single result
	if len(result) != 1 {
		t.Fatal("there was more or less than one heading")
	}

	contentWant :=
		`testtext1
testtext2
testtext3

`
	// and that results contents does not contain any of the previous text
	if result[0].Content != contentWant {
		t.Fatalf("content was not correct: want: %s, got: %s", contentWant, result[0].Content)
	}
	nameWant := "test-prefix-testfile.md"

	// and that results name matches what we see from the heading itself
	if result[0].Name != nameWant {
		t.Fatalf("name was not correct: want: %s, got: %s", nameWant, result[0].Name)
	}
}
