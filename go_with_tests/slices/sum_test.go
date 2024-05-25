package main

import "testing"

func TestSum(t *testing.T) {
	arr := [5]int{1, 2, 3, 4, 5}
	got := Sum(arr)
	want := 15

	if got != want {
		t.Errorf("got %d want %d for arr %v", got, want, arr)
	}
}
