package timecat

import "time"

func mockReadFile(f func(string) (string, error)) func(string) (string, error) {
	old := readFile
	readFile = f
	return old
}

type mockTime struct {
	time string
}

func (mt *mockTime) Format(str string) string {
	return str
}

func mockNowISODate(result string) func() string {
	old := nowISODate
	nowISODate = func() string {
		return result
	}
	return old
}

func mockReadDir(result []string) func(string) []string {
	old := readDir
	readDir = func(str string) []string {
		return result
	}
	return old
}

func mockNow(f func() time.Time) func() time.Time {
	old := now
	now = f
	return old
}
