package greetings

import "testing"

func TestHello(t *testing.T) {
	got := Hello("dark")
	want := "Hello, dark. Welcome!"

	if got != want {
		t.Errorf("Hello(\"dark\") = %q, want %q", got, want)
	}
}
