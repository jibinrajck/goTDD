package main

import "testing"

/* func TestHello(t *testing.T) {
	got := hello()
	want := "Hello World - This is my first Go Function  "

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
} */

func TestHello(t *testing.T) {

	got := hello("Raj")
	want := "Hello Raj"

	if got != want {
		t.Errorf("Want %q but got %q", want, got)
	}

}
