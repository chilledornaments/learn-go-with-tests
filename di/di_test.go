package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Test")

	got := buffer.String()
	want := "Hello, Test"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
