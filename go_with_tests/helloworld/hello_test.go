package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Test saying hello to a name", func(t *testing.T) {
		got := SayHello("Ahmad", "")
		want := "Hello, Ahmad"

		assertCorrectMessage(t, got, want)
	})

	t.Run("sending empty string greets the world", func(t *testing.T) {
		got := SayHello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := SayHello("Ahmad", "Spanish")
		want := "Hola, Ahmad"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := SayHello("Ahmad", "French")
		want := "Bonjour, Ahmad"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {

	if got != want {
		t.Errorf("got %q; want %q", got, want)
	}
}
