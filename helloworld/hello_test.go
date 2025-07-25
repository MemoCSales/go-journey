package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Memo", "")
		want := "Hello, Memo"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func (t *testing.T)  {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func (t *testing.T)  {
		got := Hello("Elodie", "French")
		want := "Bonjour, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Portuguese", func(t *testing.T) {
		got := Hello("Elodie", "Portuguese")
		want := "Olá, Elodie"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() 	// Marca esta función como una función auxiliar de prueba, no una prueba en sí misma
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}