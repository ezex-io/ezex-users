package main

import "testing"

func TestMain(t *testing.T) {
	expected := "Hello, world!"
	actual := Greet()

	if actual != expected {
		t.Errorf("Greet() = %q; want %q", actual, expected)
	}
}
