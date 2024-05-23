package main

import "testing"

func TestHello(t *testing.T) {
	got := SayHello("Ahmad")
	want := "Hello Ahmad"

	if got != want {
		t.Errorf("got %q; want %q", got, want)
	}
}
