package main

import "testing"

func TestItShouldWork(t *testing.T) {
	got := true
	want := true

	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}
