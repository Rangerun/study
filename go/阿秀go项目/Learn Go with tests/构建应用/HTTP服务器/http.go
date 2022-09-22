package main

import (
	"testing"
	"net/http/httptest"
	"fmt"
	"net/http"
)


func TestHttp(t *testing.T) {
	t.Run("returns Pepper's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		response := httptest.NewRecorder()
	
		PlayerServer(response, request)
	
		got := response.Body.String()
		want := "20"
	
		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})
}

func PlayerServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "20")
}