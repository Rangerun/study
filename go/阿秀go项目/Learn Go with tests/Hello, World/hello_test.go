package main

import "testing"

/*func TestHello(t *testing.T) {
	got := Hello("angel")
	want := "Hello, world"

	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}*/

// 子测试
/*func TestHello(t *testing.T) {

    t.Run("saying hello to people", func(t *testing.T) {
        got := Hello("Chris")
        want := "Hello, Chris"

        if got != want {
            t.Errorf("got '%q' want '%q'", got, want)
        }
    })

    t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
        got := Hello("World")
        want := "Hello, World"

        if got != want {
            t.Errorf("got '%q' want '%q'", got, want)
        }
    })

}*/

func TestHello(t *testing.T) {

    assertCorrectMessage := func(t *testing.T, got, want string) {
        t.Helper()
        if got != want {
            t.Errorf("got '%q' want '%q'", got, want)
        }
    }

    t.Run("saying hello to people", func(t *testing.T) {
        got := Hello("Chris")
        want := "Hello, Chris"
        assertCorrectMessage(t, got, want)
    })

    t.Run("empty string defaults to 'world'", func(t *testing.T) {
        got := Hello("World")
        want := "Hello, World"
        assertCorrectMessage(t, got, want)
    })

}
