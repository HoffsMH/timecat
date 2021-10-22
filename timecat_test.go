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

// when given 2021-10-21T23:35:59-05:00-hi.md it should give back the same
func TestTimestampStringWithLongValidTimestamp(t *testing.T) {
	got := TimestampString("2021-10-21T23:35:59-05:00-hi.md")
	want := "2021-10-21T23:35:59-05:00-hi.md"

	if want != got {
		t.Fatalf("want: %s, got: %s", want, got)
	}
}
