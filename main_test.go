package main

import (
	"testing"
)

func TestMatchcat(t *testing.T) {
	want := "hi"
	result := "hi"

	if result != want {
		t.Fatalf(`Ok() = %T, %T, error`, want, result)
	}
}
