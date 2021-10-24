package timecat

import (
	"testing"
	"time"
)

func TestCatWithNoFilesWhatsoever(t *testing.T) {
	var testDirContents = []string{}

	oldReadDir := mockReadDir(testDirContents)
	defer func() { readDir = oldReadDir }()

	freeze, _ := time.Parse("2006-01-02T15:04:05Z07:00", "2021-10-24T11:21:23-05:00")

	oldNow := mockNow(func() time.Time {
		return freeze
	})
	defer func() { now = oldNow }()

	oldReadFile := mockReadFile(func(f string) (string, error) {
		return "should not see this text anywhere in got", nil
	})
	defer func() { readFile = oldReadFile }()

	got := Cat("testdir", &TimeRange{0, 0, 0})
	want := "## cap.md\n"

	// no files in given directory should just have the default heading
	if got != want {
		t.Fatalf("not correct: want: %s, got: %s", want, got)
	}
}

func TestCatWithNoDatedFiles(t *testing.T) {
	var testDirContents = []string{
		"testfile",
		"testfile2",
	}

	oldReadDir := mockReadDir(testDirContents)
	defer func() { readDir = oldReadDir }()

	freeze, _ := time.Parse("2006-01-02T15:04:05Z07:00", "2021-10-24T11:21:23-05:00")

	oldNow := mockNow(func() time.Time {
		return freeze
	})
	defer func() { now = oldNow }()

	oldReadFile := mockReadFile(func(f string) (string, error) {
		return "should not see this text anywhere in got", nil
	})
	defer func() { readFile = oldReadFile }()

	got := Cat("testdir", &TimeRange{0, 0, 0})
	want := "## cap.md\n"

	// no dated files in given directory should just have the default heading
	if got != want {
		t.Fatalf("not correct: want: %s, got: %s", want, got)
	}
}

func TestCatWithOutOfRangeDate(t *testing.T) {
	var testDirContents = []string{
		"testfile",
		"testfile2",
		"2021-10-27T11:21:23-05:00-cap.md",
	}

	oldReadDir := mockReadDir(testDirContents)
	defer func() { readDir = oldReadDir }()

	freeze, _ := time.Parse("2006-01-02T15:04:05Z07:00", "2021-10-28T11:21:23-05:00")

	oldNow := mockNow(func() time.Time {
		return freeze
	})
	defer func() { now = oldNow }()

	got := Cat("testdir", &TimeRange{0, 0, 0})
	want := "## cap.md\n"

	// one dated file in the directory with SOME non dated but the date is not
	// in range
	if got != want {
		t.Fatalf("not correct: want: %s, got: %s", want, got)
	}
}

func TestCatWithOneInRange(t *testing.T) {
	var testDirContents = []string{
		"testfile",
		"testfile2",
		"2021-09-27T11:21:23-05:00-cap.md",
		"2021-10-27T11:21:23-05:00-cap.md",
	}

	oldReadDir := mockReadDir(testDirContents)
	defer func() { readDir = oldReadDir }()

	freeze, _ := time.Parse("2006-01-02T15:04:05Z07:00", "2021-10-28T11:21:23-05:00")

	oldNow := mockNow(func() time.Time {
		return freeze
	})
	defer func() { now = oldNow }()

	fileContent := `we should see this text`
	oldReadFile := mockReadFile(func(f string) (string, error) {
		return fileContent, nil
	})
	defer func() { readFile = oldReadFile }()

	got := Cat("testdir", &TimeRange{0, 0, 2})
	want := `## 2021-10-27T11:21:23-05:00-cap.md
` + fileContent + `
## cap.md
`

	// one dated file in the directory with SOME non dated but the date is not
	// in range
	if got != want {
		t.Fatalf("not correct: want: %s, got: %s", want, got)
	}
}

func TestEnsureNewline(t *testing.T) {
	got := ensureNewline("hello")
	want := "hello\n"

	if got != want {
		t.Fatalf("not correct: want: %s, got: %s", want, got)
	}

	got = ensureNewline("hello\n")

	if got != want {
		t.Fatalf("not correct: want: %s, got: %s", want, got)
	}
}
