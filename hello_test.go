package main

import "testing"

func TestHello(t *testing.T) {
    t.Run("saying hello to people", func(t *testing.T) {
    	got := Hello("Pavel")
    	want := "Hello, Pavel"
        assertCorrectMessage(t, got, want)
    })
    t.Run("saying 'Hello, World' when an empty string is supplied", func(t *testing.T) {
        got := Hello("")
        want := "Hello, World"
        assertCorrectMessage(t, got, want)
    })
}

func assertCorrectMessage(t testing.TB, got, want string) {
    t.Helper() // assure that error message will be in failing test and not here in helper function
    if got != want {
		t.Errorf("got %q want %q", got, want)
    }
}
