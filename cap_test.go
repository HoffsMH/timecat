package timecat

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)


func TestCap(t *testing.T) {
	Convey("cap usage", t, func() {
		freeze, _ := time.Parse("2006-01-02T15:04:05Z07:00", "2021-10-24T11:21:23-05:00")
		oldNow := mockNow(func() time.Time { return freeze })
		Reset(func() { now = oldNow })

		Convey("when not given a directory", func() {
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
	})
}
