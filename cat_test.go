package timecat

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)


func TestCat(t *testing.T) {
	Convey("Time is frozen at 2021-10-24T11:21:23-05:00", t, func() {
		freeze, _ := time.Parse("2006-01-02T15:04:05Z07:00", "2021-10-24T11:21:23-05:00")
		oldNow := mockNow(func() time.Time { return freeze })
		Reset(func() { now = oldNow })

		Convey("When Given an empty Directory", func() {
			var testDirContents = []string{}
			oldReadDir := mockReadDir(testDirContents)
			Reset(func() { readDir = oldReadDir })

			Convey("Should just have the cap heading", func() {
				assertCapOnly()
			})
		})

		Convey("When given a non empty directory but no dated files", func() {
			var testDirContents = []string{
				"testfile",
				"testfile2",
			}
			oldReadDir := mockReadDir(testDirContents)
			Reset(func() { readDir = oldReadDir })

			Convey("Should just have the cap heading", func() {
				assertCapOnly()
			})
		})
		Convey("When given a dir containing a dated file but date is out of range", func() {
			var testDirContents = []string{
				"testfile",
				"testfile2",
				"2021-10-23T11:21:23-05:00-cap.md",
			}
			oldReadDir := mockReadDir(testDirContents)
			Reset(func() { readDir = oldReadDir })

			Convey("Should just have the cap heading", func() {
				assertCapOnly()
			})
		})
		Convey("When given a dir containing multiple dated files and some are in the given range", func() {
			var testDirContents = []string{
				"testfile",
				"testfile2",
				"2021-09-27T11:21:23-05:00-cap.md",
				"2021-10-23T11:21:23-05:00-cap.md",
			}
			oldReadDir := mockReadDir(testDirContents)
			Reset(func() { readDir = oldReadDir })
			Convey("Should return a heading and content for the file thats in range", func() {

				fileContent := `we should see this text`
				oldReadFile := mockReadFile(func(f string) (string, error) {
					return fileContent, nil
				})
				defer func() { readFile = oldReadFile }()

				got := Cat("testdir", &TimeRange{0, 0, 2})
				want := `## 2021-10-23T11:21:23-05:00-cap.md
` + fileContent + `
## cap.md
`
				So(got, ShouldEqual, want)
			})
		})
	})
}

func TestEnsureNewline(t *testing.T) {
	Convey("when given a string that doesn't end in newline", t, func() {
		input := "hello"

		Convey("Should add a newline", func() {
				want := "hello\n"
				got := ensureNewline(input)

				So(got, ShouldEqual, want)
		})
	})
	Convey("when given a string that ends in a newline", t, func() {
		input := "hello\n"

		Convey("string shouldn't change", func() {
				got := ensureNewline(input)

				So(got, ShouldEqual, input)
		})
	})
	Convey("when given s tring that has a newline somewhere in middle", t, func() {
		input := "hel\nlo"
		Convey("Should add a newline", func() {
				want := "hel\nlo\n"
				got := ensureNewline(input)

				So(got, ShouldEqual, want)
		})
	})
}

func TestFilterFileNames(t *testing.T) {
	Convey("Given an empty array", t, func() {
		testFileNames := []string{}

		Convey("And searching for empty text", func() {
			got := filterFiles(testFileNames, "")
			So(len(got), ShouldEqual, 0)
		})

		Convey("And given non empty search text", func() {
			got := filterFiles(testFileNames, "")

			So(len(got), ShouldEqual, 0)
		})
	})

	Convey("Given a full array with some of the searched for text", t, func() {
		testFileNames := []string{
			"thisFileName",
			"2023-01-01-searchtext.md",
			"thisFileName1",
			"2021-01-01-searchtext.md",
			"thatFileName1",
		}
		Convey("And some of the text is searched for", func() {
			want := []string{
				"2023-01-01-searchtext.md",
				"2021-01-01-searchtext.md",
			}

			got := filterFiles(testFileNames, "searchtext")

			So(got, ShouldContain, want[0])
			So(got, ShouldContain, want[1])
			So(got, ShouldNotContain, testFileNames[0])
			So(got, ShouldNotContain, testFileNames[2])
			So(got, ShouldNotContain, testFileNames[4])
		})
		Convey("And none of the text is searched for", func() {
			got := filterFiles(testFileNames, "asdfg")

			for _, filename := range testFileNames {
				So(got, ShouldNotContain, filename)
			}
		})
	})
}

func assertCapOnly() {
	oldReadFile := mockReadFile(func(f string) (string, error) {
		return "should not see this text anywhere in got", nil
	})
	defer func() { readFile = oldReadFile }()

	got := Cat("testdir", &TimeRange{0, 0, 0})
	want := "## cap.md\n"

	So(got, ShouldEqual, want)
}
