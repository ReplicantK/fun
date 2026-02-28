package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to a name when a string is passed in", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"

		if got != want {
			t.Error("got:", got, "want:", want)
		}
	})

	t.Run("saying 'Hello, World' when no string is passed in", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		if got != want {
			t.Error("got:", got, "want:", want)
		}
	})
}
