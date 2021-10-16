package timecat

import (
	"testing"
)

var testFile = `
asuh
ok
`

func mockReadFile(result string) {
	readFile = func(filename string) (string, error) {
		return result, nil
	}
}

type mockTime struct {
	time string
}

func (mt *mockTime) Format(str string) string {
	return str
}

func mockNowSimpleDate(result string) {
	nowSimpleDate = func() string {
		return result
	}
}

func TestSplitWithNoHeadings(t *testing.T) {
	var testFile = `asuh
ok
asdf
asdf

`

	mockReadFile(testFile)

	result := Split("testfile")
	if len(result) > 0 {
		t.Fatal("there were headings when there should be none")
	}
}

func TestSplitWithOneHeading(t *testing.T) {
	var testFile = `asuh
ok
asdf
asdf
!# testfile.md
testtext1
testtext2
testtext3
`

	mockReadFile(testFile)

	result := Split("testfile")

	if len(result) != 1 {
		t.Fatal("there was more or less than one heading")
	}
	contentWant :=
		`testtext1
testtext2
testtext3

`
	if result[0].Content != contentWant {
		t.Fatalf("content was not correct: want: %s, got: %s", contentWant, result[0].Content)
	}
	nameWant := "testfile.md"

	if result[0].Name != nameWant {
		t.Fatalf("name was not correct: want: %s, got: %s", nameWant, result[0].Name)
	}
}

func TestSplitWithEmptyFile(t *testing.T) {
	var testFile = ``

	mockReadFile(testFile)

	result := Split("testfile")

	if len(result) != 0 {
		t.Fatal("there was more or less than one heading")
	}
}

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

	mockReadFile(testFile)

	result := Split("testfile")
	WriteSplits(result)
}

// when given hi.md it should give me back 2015-23-42-hi.md
func TestTimestampStringWithNoTimestamp(t *testing.T) {
	mockNowSimpleDate("2015-03-12")

	got := TimestampString("hi.md")
	want := "2015-03-12-hi.md"

	if want != got {
		t.Fatalf("want: %s, got: %s", want, got)
	}
}

// when given 2015-03-01-hi.md it should give me back 2015-03-01-hi.md
func TestTimstampStringWithSimpleValidTimeStamp(t *testing.T) {
	got := TimestampString("2015-03-01-hi.md")
	want := "2015-03-01-hi.md"

	if want != got {
		t.Fatalf("want: %s, got: %s", want, got)
	}
}

func TestTimstampStringWithTimestamp(t *testing.T) {
	got := TimestampString("2015-02-01-hi.md")
	want := "2015-02-01-hi.md"

	if want != got {
		t.Fatalf("want: %s, got: %s", want, got)
	}
}
