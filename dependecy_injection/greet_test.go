package greet

import (
	"bytes"
	"os"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestGreetStdOut(t *testing.T) {
	// buffer := bytes.Buffer{}
	Greet(os.Stdout, "Chris")

	// got := buffer.String()
	// want := "Hello, Chris"

	// if got != want {
	// t.Errorf("got %q want %q", got, want)
	// }
}
