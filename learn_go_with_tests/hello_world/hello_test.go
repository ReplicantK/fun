package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to a name when a string is passed in", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		if got != want {
			t.Error("got:", got, "want:", want)
		}
	})

	t.Run("saying 'Hello, World' when no string is passed in", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		if got != want {
			t.Error("got:", got, "want:", want)
		}
	})

	t.Run("testing greetings in spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		if got != want {
			t.Error("got:", got, "want:", want)
		}
	})

	t.Run("return greetings in french", func(t *testing.T) {
		got := Hello("Chris", "French")
		want := "Bonjour, Chris"

		if got != want {
			t.Error("got:", got, "want:", want)
		}
	})

	t.Run("return greetings in latin", func(t *testing.T) {
		got := Hello("Mike", "Latin")
		want := "Salve, Mike"

		if got != want {
			t.Error("got:", got, "want:", want)
		}
	})
}
