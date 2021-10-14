package timecat

import (
	"testing"
	"fmt"
)

var testFile = `
asuh
ok
`

func TestSomething(t *testing.T) {
	result := Split("testfile")
	fmt.Println(result[0].Content)
	fmt.Println(result[1].Content)
}

