package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("want '%q' got '%q'", want, got)
		}
	}
	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Glave", "")
		want := "Hello, Glave"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("qw", "Spanish")
		want := "Halo, qw"
		assertCorrectMessage(t, got, want)
	})

}
