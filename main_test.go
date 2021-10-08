package main



import (
	"testing"
)

func TestMatchcat(t *testing.T) {
	want := "hi";
	result := Ok();

	if result	 != want {
		t.Fatalf(`Ok() = %q, %v, error`, want, result)
	}
}