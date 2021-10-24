package timecat

func mockReadFile(result string, err error) func(string) (string, error) {
	oldReadFile := readFile
	readFile = func(filename string) (string, error) {
		return result, err
	}
	return oldReadFile
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
