package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectmessage := func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectmessage(t, got, want)
	})

	t.Run("saying 'Hello, world' when empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectmessage(t, got, want)
	})

	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrectmessage(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		got := Hello("Pier", "French")
		want := "Bonjour, Pier"

		assertCorrectmessage(t, got, want)
	})
}
