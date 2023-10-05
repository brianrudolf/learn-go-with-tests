package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Brian", "")
		want := "Hello, Brian"

		assertCorrectMsg(t, got, want)
	})
	t.Run("say 'Hello, world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMsg(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Brian", "Spanish")
		want := "Hola, Brian"

		assertCorrectMsg(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Brian", "French")
		want := "Bonjour, Brian"

		assertCorrectMsg(t, got, want)
	})
	t.Run("in German", func(t *testing.T) {
		got := Hello("Brian", "German")
		want := "Guten tag, Brian"

		assertCorrectMsg(t, got, want)
	})
}

func assertCorrectMsg(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

