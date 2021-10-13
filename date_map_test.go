package main

import (
	"io/fs"
	"testing"
	"time"

)

var simpleDateMaps = []*DateMap{
	&DateMap{
		mockFileInfo{"2021-08-08T14:13:45+0000-sometext.md"},
		"testdircontents/testDirEntryName",
		time.Now(),
	},
}

type mockFileInfo struct{name string}
type mockDirEntry struct{name string}

var now = time.Now()

func (s mockFileInfo) IsDir() bool        { return true }
func (s mockFileInfo) ModTime() time.Time { return now }
func (s mockFileInfo) Mode() fs.FileMode  { return 0777 }
func (s mockFileInfo) Name() string       { return s.name }
func (s mockFileInfo) Size() int64        { return 100 }
func (s mockFileInfo) Sys() interface{}   { return "asdf" }

func (s mockDirEntry) Info() (fs.FileInfo, error) { return mockFileInfo{"2021-08-08T14:13:45+0000-sometext.md"}, nil }
func (s mockDirEntry) IsDir() bool                { return true }
func (s mockDirEntry) Type() fs.FileMode          { return 0777 }

func (s mockDirEntry) Name() string {
	return s.name;
}

func stuboutfs() {
	getAbs = func(str string) (string, error) {
		return "testdircontents", nil
	}
	readDir = func(str string) ([]fs.DirEntry, error) {
		return []fs.DirEntry{mockDirEntry{"2021-08-08T14:13:45+0000-sometext.md"}}, nil
	}

	// parseDate = func(string, ...dateparse.ParserOption) (time.Time, error) {
	// 	return now, nil
	// }
}

func TestCreateDateMapsFromDir(t *testing.T) {
	stuboutfs()
	want := simpleDateMaps
	result := createDateMapsFromDir("testdir")

	if want[0].FullPath != result[0].FullPath {
		t.Fatal("FullPath doesn't match")
	}

	fstring := "2006-01-02T15:04:05.999Z"

	if want[0].Date.Format(fstring) != result[0].Date.Format(fstring) {
		t.Fatalf("Date doesn't match")
	}

	if want[0].FileInfo.Name() != result[0].FileInfo.Name() {
		t.Fatal("FileName doesn't match")
	}
}