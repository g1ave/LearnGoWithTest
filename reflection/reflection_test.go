package main

import "testing"

func TestWalk(t *testing.T) {
	expected := "Chris"

	x := struct {
		Name string
	}{expected}

	var got []string

	Walk(x, func(input string) {
		got = append(got, input)
	})

	if len(got) != 1 {
		t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
	}

	if got[0] != expected {
		t.Errorf("got %q, but want %q", got[0], expected)
	}
}
